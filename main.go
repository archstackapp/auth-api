package main

import (
	"gitlab.com/archstack/auth-api/internal/configs"
	"gitlab.com/archstack/auth-api/internal/services/auth/keycloak"
	"gitlab.com/archstack/auth-api/internal/services/users"
	"gitlab.com/archstack/workspace-api/lib/datastore"
	"gitlab.com/archstack/workspace-api/lib/logger"
)

func main() {
	configs, err := configs.NewService()
	if err != nil {
		panic(err)
	}

	logger := logger.NewService()

	keycloakConfig, err := configs.Keycloak()
	pgConfig, err := configs.Datastore()

	datastore, err := datastore.NewService(pgConfig)
	if err != nil {
		panic(err)
	}
	defer datastore.DB.Close()

	keycloak, err := keycloak.NewService(datastore, keycloakConfig)
	if err != nil {
		panic(err)
	}

	_, err := users.NewService(logger, datastore, keycloak)
	if err != nil {
		panic(err)
	}
}
