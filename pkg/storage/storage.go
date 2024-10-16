package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"gorm.io/gorm"
)

type StorageAuth interface {
	CreateUser(user *model.User) error
	GetUser(username, password string) (*model.User, error)
	DeleteUser(id int) error
}

type Storage struct {
	StorageAuth
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		NewAuth(db),
	}
}
