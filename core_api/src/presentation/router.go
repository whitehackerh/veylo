package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
)

func RegisterRoutes(
	e *echo.Echo,
	registerUserHandler *handlers.RegisterUserHandler,
	loginHandler *handlers.LoginHandler,
	getUserHandler *handlers.GetUserHandler,
	tokenManager security.TokenManager,
) {
	e.GET("/api/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	e.POST("/api/auth/register", registerUserHandler.Handle)
	e.POST("/api/auth/login", loginHandler.Handle)
	e.GET("/api/users/me", getUserHandler.Handle, handlers.AuthMiddleware(tokenManager))
}
