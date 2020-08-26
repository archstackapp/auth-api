package http

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"gitlab.com/archstack/auth-api/internal/services/auth"
	"gitlab.com/archstack/auth-api/internal/services/users"
	"gitlab.com/archstack/workspace-api/lib/server/http"
)

type Handlers struct {
	authConfig *auth.Config
	users      *users.Users
}

func NewService(authConfig *auth.Config, users *users.Users) (*Handlers, error) {
	return &Handlers{authConfig, users}, nil
}

// AddHandlers adds the echo handlers that are part of this package.
func (h *Handlers) AddHandlers(s *http.EchoServer) {
	authMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(h.authConfig.JWTSigningKey),
		SigningMethod: "RSA",
	})

	s.Echo.POST("/login", h.login)
	s.Echo.POST("/register", h.register)
	s.Echo.GET("/auth", h.checkAuth)
}

func (h *Handlers) login(c echo.Context) error {

}

func (h *Handlers) register(c echo.Context) error {

}

func (h *Handlers) checkAuth(c echo.Context) error {
	token := c.Request.Hea
}
