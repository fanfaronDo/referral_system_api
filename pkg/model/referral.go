package model

import "gorm.io/gorm"

type Referral struct {
	gorm.Model
	ReferrerId uint `json:"referrer_id"`
	ReferredId uint `json:"referred_id"`
}
