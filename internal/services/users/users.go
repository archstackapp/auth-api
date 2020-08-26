package users

import (
	uuid "github.com/gofrs/uuid"
	"gitlab.com/archstack/auth-api/internal/services/auth"
	"gitlab.com/archstack/workspace-api/lib/datastore"
	"gitlab.com/archstack/workspace-api/lib/logger"
	"gitlab.com/archstack/workspace-api/lib/models"
)

// Users struct holds all the dependencies required for the users package. And exposes all services
// provided by this package as its methods.
type Users struct {
	Repository UserRepository
	auth       auth.Auth
	logger     *logger.Logger
}

// NewService creates a new Users service.
func NewService(logger *logger.Logger, datastore *datastore.Datastore, auth auth.Auth) (*Users, error) {
	repo := UserRepository{datastore}

	w := &Users{repo, auth, logger}

	return w, nil
}

// CreateUser first creates a user on the Auth service and afterwards the user that represents the data in the apps.
func (u *Users) CreateUser(authUser *auth.User, user *models.User) (*auth.User, *models.User, error) {
	authUser, err := u.auth.CreateUser(authUser)
	if err != nil {
		return nil, nil, err
	}

	user, err = u.Repository.create(user)
	if err != nil {
		err2 := u.auth.DeleteUser(authUser)
		if err2 != nil {
			u.logger.Errorw("Unable to delete orphaned auth user",
				"err", err2.Error(),
				"authUserId", authUser.ID,
				"authUserMail", authUser.Mail)
		}

		return nil, nil, err
	}

	return authUser, user, nil
}

// GetAuthUserByMail retrieves an auth user by its mail.
func (u *Users) GetAuthUserByMail(mail string) (*auth.User, error) {
	return u.auth.GetUserByMail(mail)
}

// GetAuthUserByID retrieves an auth user by its id.
func (u *Users) GetAuthUserByID(id uuid.UUID) (*auth.User, error) {
	return u.auth.GetUserByID(id)
}

// GetDataUserByMail retrieves a data user by its mail.
func (u *Users) GetDataUserByMail(mail string) (*models.User, error) {
	return u.Repository.GetByMail(mail)
}

// GetDataUserByID retrieves a data user by its id.
func (u *Users) GetDataUserByID(id uuid.UUID) (*models.User, error) {
	return u.Repository.GetByID(id)
}
