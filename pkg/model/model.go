package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReferralCode struct {
	gorm.Model
	Code           string        `json:"code"`
	ExpirationTime time.Duration `json:"expiration_time"`
	IsActive       bool          `json:"is_active"`
	UserId         int           `json:"user_id"`
}

type Referral struct {
	gorm.Model
	ReferrerId int `json:"referrer_id"`
	ReferredId int `json:"referred_id"`
}
