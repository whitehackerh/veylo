package security

type TokenManager interface {
	Issue(userID string) (string, error)
	Verify(token string) (string, error)
}
