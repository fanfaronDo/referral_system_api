package storage

import (
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"gorm.io/gorm"
)

type StorageAuth interface {
	CreateUser(user *entry.User) error
	GetUser(username, password string) (*entry.User, error)
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
