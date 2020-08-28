package utils

import (
	"gitlab.com/archstack/auth-api/internal/services/auth"
	authmodels "gitlab.com/archstack/auth-api/lib/models"
	"gitlab.com/archstack/workspace-api/lib/models"
)

// MergeUser merges all fields of an auth.User and a models.User struct to an authmodels.FullUser one.
func MergeUser(authUser *auth.User, dataUser *models.User) *authmodels.FullUser {
	return &authmodels.FullUser{
		ID:         dataUser.ID,
		AuthID:     authUser.ID,
		FirstName:  authUser.FirstName,
		LastName:   authUser.LastName,
		Mail:       authUser.Mail,
		Level:      dataUser.Level,
		Workspaces: dataUser.Workspaces,
	}
}
