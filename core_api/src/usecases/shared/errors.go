package shared

import "errors"

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidRegistrationInput = errors.New("email is required and password must be between 8 and 72 bytes")
