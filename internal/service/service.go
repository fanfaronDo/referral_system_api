package service

import (
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"github.com/fanfaronDo/referral_system_api/internal/storage"
)

type ServiceAuth interface {
	CreateUser(user *entry.User) error
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
