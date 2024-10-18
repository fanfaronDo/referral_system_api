package service

import (
	"errors"
)

var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrUserNotFound            = errors.New("user not found")
	ErrReferrerCodeNotFound    = errors.New("referrer code not found")
	ErrReferrerCodeIsOutOfDate = errors.New("referrer code is not alive")
)
