package repositories

import "context"

type UnitOfWork interface {
	Execute(ctx context.Context, fn func(Repositories) error) error
}
