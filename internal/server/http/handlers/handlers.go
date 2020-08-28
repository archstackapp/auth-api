package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gitlab.com/archstack/auth-api/internal/server/http/middleware"
	"gitlab.com/archstack/auth-api/internal/services/auth"
	"gitlab.com/archstack/auth-api/internal/services/users"
	authmodels "gitlab.com/archstack/auth-api/lib/models"
	"gitlab.com/archstack/auth-api/lib/utils"
	"gitlab.com/archstack/workspace-api/lib/logger"
	"gitlab.com/archstack/workspace-api/lib/models"
	archhttp "gitlab.com/archstack/workspace-api/lib/server/http"
)

var validate *validator.Validate

type Handlers struct {
	auth     auth.Auth
	users    *users.Users
	logger   *logger.Logger
	validate *validator.Validate
}

func NewService(logger *logger.Logger, auth auth.Auth, users *users.Users) (*Handlers, error) {
	validate := validator.New()
	return &Handlers{auth, users, logger, validate}, nil
}

func (h *Handlers) verifyToken() {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJRVU9pRTFwMFY3WVVVcXhvdExyM2tfaThSRXlBSHpBeTVTTWxTdkZFODFnIn0.eyJleHAiOjE1OTg1NjU5MjMsImlhdCI6MTU5ODU2NTYyMywianRpIjoiYWE5ODVjYWQtYmU3OS00MTUyLThmNmItN2IyMjM1NGVkYTU1IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL2F1dGgvcmVhbG1zL3VzZXJzIiwiYXVkIjoiYWNjb3VudCIsInN1YiI6IjNkNTBhM2YyLWE2MzgtNDcxYy1hMmY0LTVlYzllMDU5MzY4NyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImF1dGgtYXBpIiwic2Vzc2lvbl9zdGF0ZSI6ImZiODczNWNkLWQyZDMtNGE5OS1hZTAwLWE4MWVjMTA1ODZjNyIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6IlBhYm9sbyBLaWxsZXJtYWNoaW5lIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiaW5mbytwYWJvbG9AbS5kZW1vbmV1bS5jb20iLCJnaXZlbl9uYW1lIjoiUGFib2xvIiwiZmFtaWx5X25hbWUiOiJLaWxsZXJtYWNoaW5lIiwiZW1haWwiOiJpbmZvK3BhYm9sb0BtLmRlbW9uZXVtLmNvbSJ9.IwgInTWjeEkcPM_hxvfKQ2rocSIYRztRHVRMSeWuWiewhh3MVn5w4s9qODYdSeOv2-qy0OW20j2NbDS9Kaiai-JOBN_Rrw5yN5OQkdEH7dFL91NtPGZOdZUNdJIJl1elE2Pe8l5jHj_UU2fxrdUbolw7_vkCIOB9WAotbGUlK4I1FQOVLNTeYY3zH87ZOMiWh0Ylmv7pBpYU2Sv24IMyhtPLflLgyyFmRbomVgMCtJPh3eTdw5eX1Xw2FnBdpiW1BLNdsktty3CTDRNFCMlhn4e18Pu8LiqRB3Nqb_Ft5rQ3wW6EA-7BRO32eOdOh1T2CYRSy5JkrN25Y9_A4Ecm4w"
	_ = "MIICmTCCAYECBgF0JvP46zANBgkqhkiG9w0BAQsFADAQMQ4wDAYDVQQDDAV1c2VyczAeFw0yMDA4MjUxODQ3MzdaFw0zMDA4MjUxODQ5MTdaMBAxDjAMBgNVBAMMBXVzZXJzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjcmJdMx2/rKFtBz955uNjvoboDBhxswSRlVfGFwf2wUBInFpUeT7KN/Q/P8WtNOabGLNvvyUi79R1tucdR4jaSqm/G51xNesDa52ukh2FgQ/4YziAt1A9AOHaCTufE3d4jl2JRHCBQDMuKIRdJu9Va01MQZ3oplrrVbeRlEWF0UZY5woNzLvuUNIl+mfuL6w5SnoS1S41EIVcbJ+i558mnsRkrlh9OLnEni/GLxXdHYJHEaA37HGpMSnlP+ruAIIRCQcqL/6ikcOe5y4EAEBy7uo+JTb8P0onSaS4wokkrB0BlJEeOA+TzgnwtHRy854Uo+rcIbuQd6V9yHas9I4eQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQBLVt7JYo4Y0D1I2bgq4VFL62lPnAkMUZItvFrYx3R5752Lm+xFcUSX8QlCKTzpnc102lGRrJsdpWW/JHnabytxNfQh8uzBoXjNSTloCfszwgazMZQ7KtBOcyY91PKyoHEq4LfHHaR6hugMmTocLJpLPqtnlUnNRvuUti9RJ/SnOC0iOS/o+iI4hAU8J3rDcu6sYBqHHF9ZA6gC37IO+nhPiDXEUXn3PLTLAfDt/3HrYuFynmb7DHU28UseofVgwd0XOUIU17oDoF9iObi8zp7SfiaahL9dqxbY0Iyyhtt3Llte6vY6DTH/tpQ7Jm0aalmn8VRndp4l5iC/Ig+UC4k2"
	signature := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjcmJdMx2/rKFtBz955uNjvoboDBhxswSRlVfGFwf2wUBInFpUeT7KN/Q/P8WtNOabGLNvvyUi79R1tucdR4jaSqm/G51xNesDa52ukh2FgQ/4YziAt1A9AOHaCTufE3d4jl2JRHCBQDMuKIRdJu9Va01MQZ3oplrrVbeRlEWF0UZY5woNzLvuUNIl+mfuL6w5SnoS1S41EIVcbJ+i558mnsRkrlh9OLnEni/GLxXdHYJHEaA37HGpMSnlP+ruAIIRCQcqL/6ikcOe5y4EAEBy7uo+JTb8P0onSaS4wokkrB0BlJEeOA+TzgnwtHRy854Uo+rcIbuQd6V9yHas9I4eQIDAQAB"

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(signature), nil
	})

	if err != nil {
		panic(err)
	}
}

// AddHandlers adds the echo handlers that are part of this package.
func (h *Handlers) AddHandlers(s *archhttp.EchoServer) {
	//h.verifyToken()
	//signingKeys := make(map[string]interface{})
	signingKey := h.auth.PublicKey()

	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    signingKey,
		SigningMethod: jwt.SigningMethodRS256.Name,
	})

	s.Echo.POST("/login", h.login)
	s.Echo.POST("/register", h.register)
	s.Echo.GET("/auth", h.checkAuth, authMiddleware)
}

// LoginRequestBody is the JSON body of a request to the login handler.
type LoginRequestBody struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (h *Handlers) login(c echo.Context) error {
	body := new(LoginRequestBody)
	if err := c.Bind(body); err != nil {
		return err
	}
	mail := body.Mail
	password := body.Password

	token, err := h.users.Login(mail, password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

// RegisterRequestBody is the JSON body of a request to the register handler.
type RegisterRequestBody struct {
	User     *authmodels.FullUser `json:"user" validate:"required"`
	Password string               `json:"password" validate:"required"`
}

func (h *Handlers) register(c echo.Context) error {
	body := new(RegisterRequestBody)
	if err := c.Bind(body); err != nil {
		return err
	}
	err := h.validate.Struct(body)
	if err != nil {
		return err
	}

	user := body.User
	h.logger.Zap.Infow("Handling registration attempt", "user", user)

	authUser := &auth.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Mail:      user.Mail,
		Password:  body.Password,
	}
	dataUser := &models.User{}

	authUser, dataUser, err = h.users.CreateUser(authUser, dataUser)
	if err != nil {
		h.logger.Zap.Debugw("Failed to create user", "user", user, "err", err)
		return err
	}

	user = utils.MergeUser(authUser, dataUser)

	return c.JSON(http.StatusOK, echo.Map{"user": user})
}

func (h *Handlers) checkAuth(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	//h.logger.Zap.Debugw("checked auth", "claims", claims)
	id := claims["sub"]

	return c.JSON(http.StatusOK, echo.Map{"id": id})
}
