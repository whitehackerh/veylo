package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	err := r.pool.QueryRow(ctx,
		`INSERT INTO users (id, email, password_hash)
		 VALUES ($1, $2, $3)
		 RETURNING created_at`,
		user.ID, user.Email, user.PasswordHash,
	).Scan(&user.CreatedAt)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == "23505" {
			return shared.ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	return r.findOne(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`, email)
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	return r.findOne(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE id = $1`, id)
}

func (r *UserRepository) findOne(ctx context.Context, query string, arg string) (*entities.User, error) {
	user := &entities.User{}
	err := r.pool.QueryRow(ctx, query, arg).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, shared.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

var _ repositories.UserRepository = (*UserRepository)(nil)
