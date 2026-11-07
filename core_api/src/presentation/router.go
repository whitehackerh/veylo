package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
)

func RegisterRoutes(e *echo.Echo, authHandlers *handlers.AuthHandlers) {
	e.GET("/api/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	e.POST("/api/auth/register", authHandlers.RegisterUser)
	e.POST("/api/auth/login", authHandlers.Login)
	e.GET("/api/users/me", authHandlers.GetUser, handlers.AuthMiddleware(authHandlers.TokenManager()))
}
