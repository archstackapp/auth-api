package keycloak

import (
	"github.com/Nerzal/gocloak/v7"
	"gitlab.com/archstack/workspace-api/lib/datastore"
)

// Keycloak implements the Auth interface.
type Keycloak struct {
	datastore *datastore.Datastore
	client    gocloak.GoCloak
	config    *Config
}

// Config struct holds all the configurations for accessing a keycloak instance.
type Config struct {
	Host      string `json:"host,omitempty"`
	Port      string `json:"port,omitempty"`
	RealmName string `json:"realmName,omitempty"`

	AdminUserName string `json:"adminUsername,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
}

// NewService creates a new Keycloak service.
func NewService(datastore *datastore.Datastore, config *Config) (*Keycloak, error) {
	client := gocloak.NewClient(config.Host + ":" + config.Port)

	w := &Keycloak{datastore, client, config}

	return w, nil
}
