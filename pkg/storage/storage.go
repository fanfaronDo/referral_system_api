package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"gorm.io/gorm"
)

type AuthStorage interface {
	CreateUser(user *model.User) error
	GetUser(username, password string) (*model.User, error)
	DeleteUser(id uint) error
}

type ReferralCodeStorage interface {
	CreateReferralCode(code *model.ReferralCode) error
	GetReferralCodeByEmail(email string) (model.ReferralCode, error)
	GetReferralCode(code string) (model.ReferralCode, error)
	GetReferralCodeByUserIdWithStatusActive(userID uint) (*model.ReferralCode, error)
	DeleteReferralCode(codeID uint) error
}

type ReferralStorage interface {
	CreateReferral(referral *model.Referral) error
	GetReferrersById(referrerId uint) ([]model.ReferralCode, error)
	GetEmailById(userId uint) (string, error)
}

type Storage struct {
	AuthStorage
	ReferralCodeStorage
	ReferralStorage
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		NewAuth(db),
		NewReferralCode(db),
		NewReferral(db),
	}
}
