package keycloak

import (
	"github.com/Nerzal/gocloak/v7"
	"github.com/gofrs/uuid"
	"gitlab.com/archstack/auth-api/internal/services/auth"
)

// NewAuthUser creates an auth.User from a gocloak.User.
func NewAuthUser(user *gocloak.User) *auth.User {
	id, err := uuid.FromString(*user.ID)
	if err != nil {
		panic(err)
	}

	return &auth.User{
		ID:        id,
		Mail:      *user.Email,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
	}
}
