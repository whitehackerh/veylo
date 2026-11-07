package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/whitehackerh/veylo/api_server/internal/domain"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(user *domain.User) error {
	err := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO users (id, email, password_hash)
		 VALUES ($1, $2, $3)
		 RETURNING created_at`,
		user.ID, user.Email, user.PasswordHash,
	).Scan(&user.CreatedAt)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == "23505" {
			return domain.ErrEmailAlreadyExists
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	return r.findOne(`SELECT id, email, password_hash, created_at FROM users WHERE email = $1`, email)
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	return r.findOne(`SELECT id, email, password_hash, created_at FROM users WHERE id = $1`, id)
}

func (r *UserRepository) findOne(query string, arg string) (*domain.User, error) {
	user := &domain.User{}
	err := r.pool.QueryRow(context.Background(), query, arg).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

var _ domain.UserRepository = (*UserRepository)(nil)
