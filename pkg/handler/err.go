package handler

import "errors"

var (
	ErrHeaderAuthUndefined    = errors.New("header auth undefined")
	ErrInvalidToken           = errors.New("invalid token")
	ErrUserNotRegistered      = errors.New("user not registered")
	ErrReferralCodeIsRequired = errors.New("referral code is required")
	ErrEmailRequired          = errors.New("email is required")
	ErrInvalidReferrerCode    = errors.New("invalid referrer code")
	ErrIncorrectId            = errors.New("incorrect id")
	ErrUserAlreadyExists      = errors.New("user already exists")
)
