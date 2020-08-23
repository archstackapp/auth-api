module gitlab.com/archstack/auth-api

go 1.14

replace gitlab.com/archstack/workspace-api => ../workspace-api

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-oauth2/oauth2 v3.9.2+incompatible
	github.com/go-oauth2/oauth2/v4 v4.1.2
	github.com/go-session/session v3.1.2+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	gitlab.com/archstack/workspace-api v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	gopkg.in/oauth2.v3 v3.12.0 // indirect
)
