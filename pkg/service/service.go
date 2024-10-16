package service

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
)

type ServiceAuth interface {
	CreateUser(user *model.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (uint, error)
}

type Service struct {
	ServiceAuth
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		NewAuth(storage),
	}
}
