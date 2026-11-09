package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/whitehackerh/veylo/api_server/src/usecases/repositories"
)

type UnitOfWork struct {
	pool *pgxpool.Pool
}

func NewUnitOfWork(pool *pgxpool.Pool) *UnitOfWork {
	return &UnitOfWork{pool: pool}
}

func (u *UnitOfWork) Execute(ctx context.Context, fn func(repositories.Repositories) error) error {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	if err := fn(NewRepositories(tx)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
