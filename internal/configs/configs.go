package configs

import (
	"time"

	"gitlab.com/archstack/auth-api/internal/services/auth"
	"gitlab.com/archstack/auth-api/internal/services/auth/keycloak"
	"gitlab.com/archstack/core-api/lib/datastore"
	"gitlab.com/archstack/core-api/lib/server/http"
)

// Configs struct handles all dependencies required for handling configurations
type Configs struct {
}

// HTTP returns the configuration required for HTTP package
func (cfg *Configs) HTTP() (*http.Config, error) {
	return &http.Config{
		Port:         "3333",
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
	}, nil
}

// Keycloak returns the configuration required for the keycloak auth impl.
func (cfg *Configs) Keycloak() (*keycloak.Config, error) {
	return &keycloak.Config{
		Host:          "http://localhost",
		Port:          "8080",
		ClientID:      "auth-api",
		ClientSecret:  "0a710d2f-7978-4a38-b560-db7eb476d2d0",
		RealmName:     "users",
		AdminUserName: "admin",
		AdminPassword: "admin",
		Kid:           "QUOiE1p0V7YUUqxotLr3k_i8REyAHzAy5SMlSvFE81g",
	}, nil
}

// Auth returns the configuration for the general auth package.
func (cfg *Configs) Auth() (*auth.Config, error) {
	return &auth.Config{
		JWTSigningKey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjcmJdMx2/rKFtBz955uNjvoboDBhxswSRlVfGFwf2wUBInFpUeT7KN/Q/P8WtNOabGLNvvyUi79R1tucdR4jaSqm/G51xNesDa52ukh2FgQ/4YziAt1A9AOHaCTufE3d4jl2JRHCBQDMuKIRdJu9Va01MQZ3oplrrVbeRlEWF0UZY5woNzLvuUNIl+mfuL6w5SnoS1S41EIVcbJ+i558mnsRkrlh9OLnEni/GLxXdHYJHEaA37HGpMSnlP+ruAIIRCQcqL/6ikcOe5y4EAEBy7uo+JTb8P0onSaS4wokkrB0BlJEeOA+TzgnwtHRy854Uo+rcIbuQd6V9yHas9I4eQIDAQAB",
		Kid:           "QUOiE1p0V7YUUqxotLr3k_i8REyAHzAy5SMlSvFE81g",
	}, nil
}

// Datastore returns datastore configuration
func (cfg *Configs) Datastore() (*datastore.Config, error) {
	return &datastore.Config{
		Host:   "localhost",
		Port:   "5432",
		Driver: "postgres",

		StoreName: "archstack",
		Username:  "",
		Password:  "",

		SSLMode: "",

		ConnPoolSize: 10,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 60,
		DialTimeout:  time.Second * 10,
	}, nil
}

// NewService returns an instance of Config with all the required dependencies initialized
func NewService() (*Configs, error) {
	return &Configs{}, nil
}
