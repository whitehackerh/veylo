package security

type TokenManager interface {
	Issue(userId string) (string, error)
	Verify(token string) (string, error)
}
