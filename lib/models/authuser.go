package models

import (
	"github.com/gofrs/uuid"
	"gitlab.com/archstack/workspace-api/lib/models"
)

// AuthenticatedUser represents an authenticated user of the apps
type AuthenticatedUser struct {
	ID     uuid.UUID `json:"id"`
	AuthID uuid.UUID `json:"-"`

	Mail      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`

	Level int `json:"level" pg:",notnull,default:0"`

	Workspaces []*models.Workspace `json:"workspaces"`

	InvitedBy *models.User `json:"invitedBy"`
}
