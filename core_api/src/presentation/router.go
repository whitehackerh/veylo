package presentation

import (
	"net/http"

	"github.com/labstack/echo/v5"
	echomiddleware "github.com/labstack/echo/v5/middleware"

	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/presentation/middlewares"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
)

type Handlers struct {
	SignUp *handlers.SignUpHandler
	SignIn *handlers.SignInHandler
	GetUser *handlers.GetUserHandler
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
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8082",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		ExposeHeaders: []string{echo.HeaderAuthorization},
	}))

	api := e.Group("/api")

	api.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, presenters.ToAPIResponse(c.Request().URL.Path, presenters.ToHealthData(), nil))
	})
	api.POST("/sign-up", handlers.SignUp.Handle)
	api.POST("/sign-in", handlers.SignIn.Handle)

	authenticated := api.Group("")
	authenticated.Use(middlewares.Auth(tokenManager, useCases.GetUser))

	authenticated.GET("/users/me", handlers.GetUser.Handle)
}
