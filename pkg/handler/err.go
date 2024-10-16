package handler

import "errors"

var (
	ErrHeaderAuthUndefined = errors.New("header auth undefined")
	ErrInvalidToken        = errors.New("invalid token")
)