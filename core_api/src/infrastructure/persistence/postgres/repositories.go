package postgres

import "github.com/whitehackerh/veylo/api_server/src/usecases/repositories"

type RepositorySet struct {
	userRepository repositories.UserRepository
}

func NewRepositories(db DBTX) *RepositorySet {
	return &RepositorySet{
		userRepository: NewUserRepository(db),
	}
}

func (r *RepositorySet) User() repositories.UserRepository {
	return r.userRepository
}
