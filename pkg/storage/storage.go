package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"gorm.io/gorm"
)

type AuthStorage interface {
	CreateUser(user *model.User) error
	GetUser(username, password string) (*model.User, error)
	DeleteUser(id int) error
}

type ReferralStorage interface {
	CreateReferral(referral *model.Referral) error
	CreateReferralCode(code *model.ReferralCode) error
	GetReferralCodeByEmail(email string) (model.ReferralCode, error)
	GetReferrersById(referrerId int) ([]model.ReferralCode, error)
	GetReferralCode(code string) (model.ReferralCode, error)
	GetEmailById(userId int) (string, error)
	DeleteReferralCode(codeID uint) error
}

type Storage struct {
	AuthStorage
	ReferralStorage
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		NewAuth(db),
		NewReferral(db),
	}
}
