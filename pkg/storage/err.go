package storage

import "errors"

var (
	ErrUserCodeNotFound           = errors.New("user code not found")
	ErrActiveReferralCodeNotFound = errors.New("active referral code not found")
	ErrReferralCodeNotFound       = errors.New("referral code not found")
)
