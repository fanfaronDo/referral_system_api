package service

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
)

const (
	MaxLengthReferralCode = 8
)

type AuthService interface {
	CreateUser(user *model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (uint, error)
	IsUserExists(username string) bool
}

type ReferralCodeService interface {
	CreateReferralCode(referralCode *model.ReferralCode) (string, error)
	GetReferralCodeByEmail(userID uint, email string) (model.ReferralCode, error)
	GetReferralCode(code string) (*model.ReferralCode, error)
	CheckReferralCode(referralCode *model.ReferralCode) error
	UpdateReferralCodeStatus(referralCode *model.ReferralCode, status bool) error
	DeleteReferralCode(userID uint, code string) error
}

type ReferralService interface {
	CreateReferral(referralCode *model.ReferralCode, userID uint) error
	GetReferrersById(referrerId uint) ([]model.ReferralCode, error)
}

type Service struct {
	AuthService
	ReferralCodeService
	ReferralService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		NewAuth(storage),
		NewReferralCode(storage),
		NewReferral(storage),
	}
}
