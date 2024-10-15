package storage

import (
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (auth *Auth) CreateUser(user *entry.User) error {
	tx := auth.db.Create(user)
	return tx.Error
}

func (auth *Auth) GetUser(username, password string) (*entry.User, error) {
	var user entry.User
	auth.db.Where("username = ? AND password = ?", username, password).First(&user)
	return &user, nil
}

func (auth *Auth) DeleteUser(id int) error {
	tx := auth.db.Where("id = ?", id).Delete(&entry.User{})
	return tx.Error
}
