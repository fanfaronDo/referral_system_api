package model

import "time"

type ReferralInfo struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}
