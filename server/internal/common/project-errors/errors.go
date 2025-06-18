package projectErrors

import "errors"

var (
	ErrIllegalPassword = errors.New("password given is invalid")
	ErrUserNotUnique   = errors.New("users already exists")
)
