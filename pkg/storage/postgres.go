package storage

import (
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(host string, config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host,
			config.User,
			config.Password,
			config.Database,
			config.Postgres.Port,
			config.SSLMode),
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
