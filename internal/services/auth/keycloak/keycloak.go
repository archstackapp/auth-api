package keycloak

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"github.com/Nerzal/gocloak/v7"
	"gitlab.com/archstack/auth-api/internal/services/auth"
	"gitlab.com/archstack/workspace-api/lib/datastore"
)

// Keycloak implements the Auth interface.
type Keycloak struct {
	datastore        *datastore.Datastore
	client           gocloak.GoCloak
	config           *Config
	token            *gocloak.JWT
	tokenLastUpdated time.Time
}

// Config struct holds all the configurations for accessing a keycloak instance.
type Config struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`

	ClientID     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`

	RealmName string `json:"realmName,omitempty"`

	AdminUserName string `json:"adminUsername,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
}

// NewService creates a new Keycloak service.
func NewService(datastore *datastore.Datastore, config *Config) (*Keycloak, error) {
	client := gocloak.NewClient(config.Host + ":" + config.Port)

	initialToken, err := getToken(context.Background(), config, client)
	if err != nil {
		panic(err)
	}

	w := &Keycloak{
		datastore:        datastore,
		client:           client,
		config:           config,
		token:            initialToken,
		tokenLastUpdated: time.Now(),
	}

	return w, nil
}

func getToken(ctx context.Context, config *Config, client gocloak.GoCloak) (*gocloak.JWT, error) {
	token, err := client.LoginAdmin(ctx, config.AdminUserName, config.AdminPassword, config.RealmName)

	return token, err
}

func (k *Keycloak) getToken(ctx context.Context) (*gocloak.JWT, error) {
	now := time.Now()
	if now.After(k.tokenLastUpdated.Add(time.Hour * time.Duration(3))) {
		newToken, err := getToken(ctx, k.config, k.client)
		if err != nil {
			return nil, err
		}

		k.token = newToken
		k.tokenLastUpdated = now
	}

	return k.token, nil
}

// CreateUser creates a new user on the keycloak instance.
func (k *Keycloak) CreateUser(user *auth.User) (*auth.User, error) {
	ctx := context.Background()
	token, err := k.getToken(ctx)
	if err != nil {
		return nil, err
	}

	enabled := true

	keycloakUser := gocloak.User{
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Email:     &user.Mail,
		Enabled:   &enabled,
		Username:  &user.Mail,
	}

	idString, err := k.client.CreateUser(ctx, token.AccessToken, k.config.RealmName, keycloakUser)

	id, err := uuid.FromString(idString)
	if err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}

// GetUserByMail retrieves a user from the keycloak instance by its mail.
func (k *Keycloak) GetUserByMail(mail string) (*auth.User, error) {
	ctx := context.Background()
	token, err := k.getToken(ctx)
	if err != nil {
		return nil, err
	}

	params := gocloak.GetUsersParams{Email: &mail}
	users, err := k.client.GetUsers(ctx, token.AccessToken, k.config.RealmName, params)
	if err != nil {
		return nil, err
	}

	if len(users) > 1 {
		panic(errors.New("keycloak should only return 1 user for a mail"))
	}

	return NewAuthUser(users[0]), nil
}

// GetUserByID retrieves a user from the keycloak instance by its id.
func (k *Keycloak) GetUserByID(id uuid.UUID) (*auth.User, error) {
	ctx := context.Background()
	token, err := k.getToken(ctx)
	if err != nil {
		return nil, err
	}

	user, err := k.client.GetUserByID(ctx, token.AccessToken, k.config.RealmName, id.String())
	if err != nil {
		return nil, err
	}

	return NewAuthUser(user), nil
}

// DeleteUser deletes a user from the keycloak instance by its id.
func (k *Keycloak) DeleteUser(id uuid.UUID) error {
	ctx := context.Background()
	token, err := k.getToken(ctx)
	if err != nil {
		return err
	}

	err = k.client.DeleteUser(ctx, token.AccessToken, k.config.RealmName, id.String())
	if err != nil {
		return err
	}

	return nil
}

func (k *Keycloak) Login(mail string, password string) (string, error) {
	ctx := context.Background()

	token, err := k.client.Login(ctx, k.config.ClientID, k.config.ClientSecret, k.config.RealmName, mail, password)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (k *Keycloak) CheckToken(token string) error {
	const tokenNotActiveMessage = "token isn't active"

	ctx := context.Background()

	tokenValidation, _, err := k.client.DecodeAccessToken(ctx, token, k.config.RealmName, "")
	if err != nil {
		return err
	}
	println("foudn tokenvalidation with header ", tokenValidation.Header)

	// rptResult, err := k.client.RetrospectToken(ctx, token, k.config.ClientID, k.config.ClientSecret, k.config.RealmName)
	// if err != nil {
	// 	return err
	// }

	if !tokenValidation.Valid {
		return errors.New(tokenNotActiveMessage)
	}

	//permissions := rptResult.Permissions
	return nil
}

func (k *Keycloak) RetrieveAuthUserFromToken(token string) (*auth.User, error) {
	ctx := context.Background()

	userInfo, err := k.client.GetUserInfo(ctx, token, k.config.RealmName)

	id, err := uuid.FromString(*userInfo.Sub)
	if err != nil {
		return nil, err
	}

	user := &auth.User{
		ID:        id,
		FirstName: *userInfo.GivenName,
		LastName:  *userInfo.FamilyName,
		Mail:      *userInfo.Email,
	}

	return user, nil
}
