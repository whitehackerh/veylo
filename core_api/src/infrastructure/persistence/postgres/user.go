package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/whitehackerh/veylo/api_server/src/domain/entities"
	"github.com/whitehackerh/veylo/api_server/src/usecases/shared"
)

type UserRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	err := r.db.QueryRow(ctx,
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
	var id string
	var email string
	var passwordHash string
	var createdAt time.Time
	err := r.db.QueryRow(ctx, query, arg).Scan(&id, &email, &passwordHash, &createdAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, shared.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return entities.RestoreUser(id, email, passwordHash, createdAt), nil
}
