package bootstrap

import "errors"

var ErrInvalidConfiguration = errors.New("DATABASE_URL and a JWT_SECRET of at least 32 characters must be set")
