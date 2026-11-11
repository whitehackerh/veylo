package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/presentation/middlewares"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
)

type Handlers struct {
	RegisterUser *handlers.RegisterUserHandler
	Login        *handlers.LoginHandler
	GetUser      *handlers.GetUserHandler
}

type UseCases struct {
	GetUser ports.UseCase[dtos.GetUserInput, dtos.GetUserOutput]
}

func RegisterRoutes(
	e *echo.Echo,
	handlers Handlers,
	useCases UseCases,
	tokenManager security.TokenManager,
) {
	api := e.Group("/api")

	api.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, presenters.ToAPIResponse(c.Request().URL.Path, presenters.ToHealthData(), nil))
	})
	api.POST("/register", handlers.RegisterUser.Handle)
	api.POST("/login", handlers.Login.Handle)

	authenticated := api.Group("")
	authenticated.Use(middlewares.Auth(tokenManager, useCases.GetUser))

	authenticated.GET("/users/me", handlers.GetUser.Handle)
}
