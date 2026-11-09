package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
)

func RegisterRoutes(
	e *echo.Echo,
	registerUserHandler *handlers.RegisterUserHandler,
	loginHandler *handlers.LoginHandler,
	getUserHandler *handlers.GetUserHandler,
	tokenManager security.TokenManager,
	getUserUseCase ports.GetUserUseCase,
) {
	api := e.Group("/api")

	api.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	api.POST("/auth/register", registerUserHandler.Handle)
	api.POST("/auth/login", loginHandler.Handle)
	api.GET("/users/me", getUserHandler.Handle, handlers.AuthMiddleware(tokenManager, getUserUseCase))
}
