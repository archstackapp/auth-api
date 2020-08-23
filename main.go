package main

import (
	"gitlab.com/archstack/auth-api/configs"
	"gitlab.com/archstack/auth-api/internal/server/oauth2"
	"gitlab.com/archstack/auth-api/internal/services/users"
	"gitlab.com/archstack/workspace-api/lib/datastore"
)

func main() {
	configs, err := configs.NewService()
	if err != nil {
		panic(err)
	}

	oauth2Config, err := configs.OAuth2()
	pgConfig, err := configs.Datastore()

	datastore, err := datastore.NewService(pgConfig)
	if err != nil {
		panic(err)
	}
	defer datastore.DB.Close()

	users, err := users.NewService(datastore)

	oauth2, err := oauth2.NewService(oauth2Config, users)
	if err != nil {
		panic(err)
	}

	oauth2.Start()
}
