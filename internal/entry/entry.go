package entry

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Referral struct {
	gorm.Model
	Code           string        `json:"code"`
	ExpirationTime time.Duration `json:"expiration_time"`
	UserId         int           `json:"user_id"`
}
