package projectErrors

import "errors"

var (
	ErrIllegalPasswordError = errors.New("password given is invalid")
	ErrUserNotUnique        = errors.New("users already exists")
)
