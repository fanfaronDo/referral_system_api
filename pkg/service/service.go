package service

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
)

type AuthService interface {
	CreateUser(user *model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (uint, error)
}

type ReferralService interface {
	CreateReferralCode(referralCode *model.ReferralCode) (string, error)
	GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error)
	DeleteReferralCode(userID uint, code string) error
}

type Service struct {
	AuthService
	ReferralService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		NewAuth(storage),
		NewReferral(storage),
	}
}
