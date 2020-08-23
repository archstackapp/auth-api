package users

import (
	"gitlab.com/archstack/workspace-api/lib/datastore"
)

// Users struct holds all the dependencies required for the users package. And exposes all services
// provided by this package as its methods
type Users struct {
	Repository UserRepository
}

// NewService creates a new Users service
func NewService(datastore *datastore.Datastore) (*Users, error) {
	repo := UserRepository{datastore}

	w := &Users{repo}

	return w, nil
}
