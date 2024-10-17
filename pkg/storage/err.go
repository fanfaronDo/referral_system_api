package storage

import "errors"

var (
	ErrUserCodeNotFound           = errors.New("user code not found")
	ErrActiveReferralCodeNotFound = errors.New("active referral code not found")
)
