package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"

	"github.com/whitehackerh/veylo/api_server/src/infrastructure/persistence/postgres"
	"github.com/whitehackerh/veylo/api_server/src/infrastructure/security"
	"github.com/whitehackerh/veylo/api_server/src/infrastructure/storages"
	"github.com/whitehackerh/veylo/api_server/src/presentation"
	"github.com/whitehackerh/veylo/api_server/src/presentation/handlers"
	"github.com/whitehackerh/veylo/api_server/src/presentation/presenters"
	"github.com/whitehackerh/veylo/api_server/src/usecases/interactors"
)

type App struct {
	Server *echo.Echo
	Pool   *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	storagePath := os.Getenv("STORAGE_PATH")
	mediaBaseUrl := os.Getenv("MEDIA_BASE_URL")
	if databaseURL == "" || len(jwtSecret) < 32 || storagePath == "" || mediaBaseUrl == "" {
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

	tokenManager := security.NewJWTManager(jwtSecret, "veylo-core-api", 15*time.Minute)

	imageStorage := storages.NewLocalImageStorage(storagePath)

	uow := postgres.NewUnitOfWork(pool)
	userRepository := postgres.NewUserRepository(pool)
	// channelRepository := postgres.NewChannelRepository(pool)

	signUpUseCase := interactors.NewSignUpInteractor(userRepository)
	signInUseCase := interactors.NewSignInInteractor(userRepository, tokenManager)
	getUserUseCase := interactors.NewGetUserInteractor(userRepository)
	createChannelUseCase := interactors.NewCreateChannelInteractor(imageStorage, uow)

	responsePresenter := presenters.NewResponsePresenter()
	channelPresenter := presenters.NewChannelPresenter(mediaBaseUrl)
	userPresenter := presenters.NewUserPresenter()
	createChannelPresenter := presenters.NewCreateChannelPresenter(channelPresenter)
	getUserPresenter := presenters.NewGetUserPresenter(userPresenter)
	signInPresenter := presenters.NewSignInPresenter()
	signUpPresenter := presenters.NewSignUpPresenter()

	server := echo.New()
	server.Static("/media/images", filepath.Join(storagePath, "images"))
	presentation.RegisterRoutes(
		server,
		presentation.Handlers{
			CreateChannel: handlers.NewCreateChannelHandler(responsePresenter, createChannelUseCase, createChannelPresenter),
			GetUser:       handlers.NewGetUserHandler(responsePresenter, getUserPresenter),
			SignIn:        handlers.NewSignInHandler(responsePresenter, signInUseCase, signInPresenter),
			SignUp:        handlers.NewSignUpHandler(responsePresenter, signUpUseCase, signUpPresenter),
		},
		responsePresenter,
		getUserUseCase,
		tokenManager,
	)

	return &App{Server: server, Pool: pool}, nil
}
