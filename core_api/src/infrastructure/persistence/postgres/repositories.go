package postgres

import "github.com/whitehackerh/veylo/api_server/src/usecases/repositories"

type RepositorySet struct {
	userRepository    repositories.UserRepository
	channelRepository repositories.ChannelRepository
}

func NewRepositories(db DBTX) *RepositorySet {
	return &RepositorySet{
		userRepository:    NewUserRepository(db),
		channelRepository: NewChannelRepository(db),
	}
}

func (r *RepositorySet) User() repositories.UserRepository {
	return r.userRepository
}

func (r *RepositorySet) Channel() repositories.ChannelRepository {
	return r.channelRepository
}
