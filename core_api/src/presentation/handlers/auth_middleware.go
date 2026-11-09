package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/presentation/openapi"
	"github.com/whitehackerh/veylo/api_server/src/usecases/dtos"
	"github.com/whitehackerh/veylo/api_server/src/usecases/ports"
	"github.com/whitehackerh/veylo/api_server/src/usecases/security"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

func AuthMiddleware(tokenManager security.TokenManager, getUserUseCase ports.GetUserUseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			const prefix = "Bearer "
			if !strings.HasPrefix(authorization, prefix) {
				return c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{Error: "unauthorized"})
			}

			userID, err := tokenManager.Verify(strings.TrimPrefix(authorization, prefix))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{Error: "unauthorized"})
			}

			output, err := getUserUseCase.Execute(c.Request().Context(), dtos.GetUserInput{UserID: userID})
			if err != nil {
				if errors.Is(err, shared.ErrUserNotFound) {
					return c.JSON(http.StatusUnauthorized, openapi.ErrorResponse{Error: "unauthorized"})
				}
				return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Error: "failed to authenticate user"})
			}

			c.Set("user", output.User)
			return next(c)
		}
	}
}
