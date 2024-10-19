package model

import (
	"gorm.io/gorm"
	"time"
)

type ReferralCode struct {
	gorm.Model
	Code           string        `json:"code" gorm:"column:code"`
	ExpirationTime time.Duration `json:"expiration_time" gorm:"column:expiration_time"`
	IsActive       bool          `json:"is_active" gorm:"column:is_active"`
	UserId         uint          `json:"user_id" gorm:"column:user_id"`
}

func (r *ReferralCode) UpdateAliveTimeStatus() {
	currentTime := time.Now()
	timealive := currentTime.Sub(r.CreatedAt)
	if timealive > r.ExpirationTime {
		r.IsActive = false
	} else {
		r.IsActive = true
	}
}
