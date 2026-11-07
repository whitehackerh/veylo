package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
)

func NewRouter(authHandler *handlers.AuthHandler, tokenManager security.TokenManager) *echo.Echo {
	e := echo.New()
	e.GET("/api/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	e.POST("/api/auth/register", authHandler.Register)
	e.POST("/api/auth/login", authHandler.Login)
	e.GET("/api/users/me", authHandler.Me, handlers.AuthMiddleware(tokenManager))
	return e
}
