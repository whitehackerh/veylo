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
	CreateChannel *handlers.CreateChannelHandler
	GetUser       *handlers.GetUserHandler
	SignIn        *handlers.SignInHandler
	SignUp        *handlers.SignUpHandler
}

func RegisterRoutes(
	e *echo.Echo,
	handlers Handlers,
	responsePresenter *presenters.ResponsePresenter,
	getUserUseCase ports.UseCase[dtos.GetUserInput, dtos.GetUserOutput],
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
	api.POST("/sign-in", handlers.SignIn.Handle)
	api.POST("/sign-up", handlers.SignUp.Handle)

	authenticated := api.Group("")
	authenticated.Use(middlewares.Auth(tokenManager, responsePresenter, getUserUseCase))

	authenticated.GET("/users/me", handlers.GetUser.Handle)
	authenticated.POST("/channels", handlers.CreateChannel.Handle)
}
