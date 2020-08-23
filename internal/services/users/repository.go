package users

import (
	uuid "github.com/gofrs/uuid"
	"gitlab.com/archstack/workspace-api/lib/datastore"
	"gitlab.com/archstack/workspace-api/lib/models"
)

// UserRepository enables CRUD ops on the db for the User objects.
type UserRepository struct {
	datastore *datastore.Datastore
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := new(models.User)

	err := r.datastore.DB.Model(user).
		Where("id = ?", id).
		Select()

	return user, err
}

func (r *UserRepository) GetByMail(mail string) (*models.User, error) {
	user := new(models.User)

	err := r.datastore.DB.Model(user).
		Where("mail = ?", mail).
		Select()

	return user, err
}
