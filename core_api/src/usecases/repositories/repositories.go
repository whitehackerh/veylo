package repositories

type Repositories interface {
	User() UserRepository
	Channel() ChannelRepository
}
