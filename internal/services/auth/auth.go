package auth

import "github.com/gofrs/uuid"

// Auth struct holds all the dependencies required for the auth package. And exposes all services
// provided by this package as its methods.
type Auth interface {
}

// User represents the data that is necessary to authenticate a user with the apps.
type User interface {
	ID() uuid.UUID
	FirstName() string
	LastName() string
	Mail() string
}
