module gitlab.com/archstack/auth-api

go 1.14

replace gitlab.com/archstack/core-api => ../core-api

require (
	github.com/Nerzal/gocloak v1.0.0
	github.com/Nerzal/gocloak/v7 v7.4.0
	github.com/dghubble/sling v1.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/go-oauth2/oauth2 v3.9.2+incompatible
	github.com/go-oauth2/oauth2/v4 v4.1.2
	github.com/go-pg/pg/v10 v10.7.6
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-session/session v3.1.2+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gorilla/sessions v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.2.0
	github.com/pkg/errors v0.9.1
	github.com/segmentio/encoding v0.1.15 // indirect
	gitlab.com/archstack/core-api v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/exp v0.0.0-20200821190819-94841d0725da // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	gopkg.in/oauth2.v3 v3.12.0 // indirect
)
