package storage

import (
	"github.com/fanfaronDo/referral_system_api/pkg/model"
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

func (auth *Auth) CreateUser(user *model.User) error {
	tx := auth.db.Create(user)
	return tx.Error
}

func (auth *Auth) GetUser(username, password string) (*model.User, error) {
	var user model.User
	tx := auth.db.Where("username = ? AND password = ?", username, password).First(&user)
	return &user, tx.Error
}

func (auth *Auth) DeleteUser(id uint) error {
	tx := auth.db.Where("id = ?", id).Delete(&model.User{})
	return tx.Error
}
