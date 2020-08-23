package oauth2

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"gitlab.com/archstack/auth-api/internal/services/users"
)

type OAuth2 struct {
	config       *Config
	manager      *manage.Manager
	clientStore  *store.ClientStore
	oauth2Server *server.Server
	users        *users.Users
}

func NewService(config *Config, users *users.Users) (*OAuth2, error) {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: config.ClientDomainURL,
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)

	oauth2 := &OAuth2{
		config,
		manager,
		clientStore,
		srv,
		users,
	}

	srv.SetPasswordAuthorizationHandler(oauth2.passwordAuthorizationHandler)

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return oauth2, nil
}

// Start the OAuth2 server and listen on the port specified in the config.
func (o *OAuth2) Start() {
	log.Fatal(http.ListenAndServe(":"+o.config.Port, nil))
}
