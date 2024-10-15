package app

import (
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"github.com/fanfaronDo/referral_system_api/internal/storage"
	"github.com/fanfaronDo/referral_system_api/migrations"
)

func Run(cnf *config.Config) error {
	db, err := storage.NewPostgres(cnf.Postgres.Host, cnf)
	if err != nil {
		return err
	}
	migrator := migrations.NewMigrator(db)
	err = migrator.MigrateUp(&entry.User{}, &entry.Referral{})

	if err != nil {
		fmt.Println(err)
	}

	auth := storage.NewAuth(db)

	//user := entry.User{
	//	Username: "sharipsvs@mail.ru",
	//	Password: "password2",
	//}

	e := auth.DeleteUser(2)
	fmt.Println(e)

	return nil
}
