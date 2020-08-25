package keycloak

import (
	"github.com/Nerzal/gocloak"
	"github.com/gofrs/uuid"
)

// User wraps a gocloak.User and implements the User interface from the auth package.
type User struct {
	inner *gocloak.User
}

// ID returns the id of the user.
func (u *User) ID() uuid.UUID {
	id, err := uuid.FromString(u.inner.ID)
	if err != nil {
		panic(err)
	}
	return id
}

// FirstName returns the first name of the user.
func (u *User) FirstName() string {
	return u.inner.FirstName
}

// LastName returns the last name of the user.
func (u *User) LastName() string {
	return u.inner.LastName
}

// Mail returns the mail of the user.
func (u *User) Mail() string {
	return u.inner.Email
}
