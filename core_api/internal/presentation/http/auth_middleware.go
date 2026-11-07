package http

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/internal/usecase"
)

func AuthMiddleware(tokenManager usecase.TokenManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			const prefix = "Bearer "
			if !strings.HasPrefix(authorization, prefix) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			}
			userID, err := tokenManager.Verify(strings.TrimPrefix(authorization, prefix))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			}
			c.Set("user_id", userID)
			return next(c)
		}
	}
}
