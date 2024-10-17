package service

import (
	"errors"
)

var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrUserNotFound            = errors.New("user not found")
)
