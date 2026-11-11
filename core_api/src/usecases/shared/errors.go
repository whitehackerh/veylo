package shared

import "errors"

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidSignUpInput = errors.New("invalid email or password")
