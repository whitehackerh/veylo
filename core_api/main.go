package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/internal/infrastructure/postgres"
	"github.com/whitehackerh/veylo/api_server/internal/infrastructure/security"
	httpdelivery "github.com/whitehackerh/veylo/api_server/internal/presentation/http"
	"github.com/whitehackerh/veylo/api_server/internal/usecase"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if databaseURL == "" || len(jwtSecret) < 32 {
		log.Fatal("DATABASE_URL and a JWT_SECRET of at least 32 characters must be set")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	userRepository := postgres.NewUserRepository(pool)
	tokenManager := security.NewJWTManager(jwtSecret, "veylo-core-api", 15*time.Minute)
	authUsecase := usecase.NewAuthUsecase(userRepository, tokenManager)
	authHandler := httpdelivery.NewAuthHandler(authUsecase)

	e := echo.New()
	e.GET("/api/health", func(c *echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	e.POST("/api/auth/register", authHandler.Register)
	e.POST("/api/auth/login", authHandler.Login)
	e.GET("/api/users/me", authHandler.Me, httpdelivery.AuthMiddleware(tokenManager))

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API server listening on :%s", port)
	log.Fatal(e.Start(":" + port))
}
