package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/infrastructure/repositories/postgres"
	"github.com/whitehackerh/veylo/api_server/src/infrastructure/security"
	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/usecases/interactors"
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
	registerUserInteractor := interactors.NewRegisterUserInteractor(userRepository)
	loginInteractor := interactors.NewLoginInteractor(userRepository, tokenManager)
	getUserInteractor := interactors.NewGetUserInteractor(userRepository)
	authHandler := handlers.NewAuthHandler(registerUserInteractor, loginInteractor, getUserInteractor)

	e := echo.New()
	e.GET("/api/health", func(c *echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	e.POST("/api/auth/register", authHandler.Register)
	e.POST("/api/auth/login", authHandler.Login)
	e.GET("/api/users/me", authHandler.Me, handlers.AuthMiddleware(tokenManager))

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API server listening on :%s", port)
	log.Fatal(e.Start(":" + port))
}
