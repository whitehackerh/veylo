package bootstrap

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/infrastructure/persistence/postgres"
	"github.com/whitehackerh/veylo/api_server/src/infrastructure/security"
	"github.com/whitehackerh/veylo/api_server/src/presentation"
	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/usecases/interactors"
)

type App struct {
	Server *echo.Echo
	Pool   *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if databaseURL == "" || len(jwtSecret) < 32 {
		return nil, ErrInvalidConfiguration
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	userRepository := postgres.NewUserRepository(pool)
	tokenManager := security.NewJWTManager(jwtSecret, "veylo-core-api", 15*time.Minute)

	registerUserUseCase := interactors.NewRegisterUserInteractor(userRepository)
	loginUseCase := interactors.NewLoginInteractor(userRepository, tokenManager)
	getUserUseCase := interactors.NewGetUserInteractor(userRepository)

	server := echo.New()
	presentation.RegisterRoutes(
		server,
		presentation.Handlers{
			RegisterUser: handlers.NewRegisterUserHandler(registerUserUseCase),
			Login:        handlers.NewLoginHandler(loginUseCase),
			GetUser:      handlers.NewGetUserHandler(),
		},
		presentation.UseCases{
			GetUser: getUserUseCase,
		},
		tokenManager,
	)

	return &App{Server: server, Pool: pool}, nil
}
