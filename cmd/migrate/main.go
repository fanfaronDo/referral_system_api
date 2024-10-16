package main

import (
	"flag"
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"github.com/fanfaronDo/referral_system_api/migrations"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
	"log"
)

func main() {

	flagMigrateUp := flag.Bool("up", false, "usage command with [migrate up]")
	flagMigrateDown := flag.Bool("down", false, "usage command with [migrate down]")
	flag.Parse()

	if !(*flagMigrateUp || *flagMigrateDown) {
		log.Fatalf("usage command with [migrate --up|down]")
		return
	}

	cnf := config.ConfigLoad()
	db, err := storage.NewPostgres(cnf.Postgres.Host, cnf)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}

	migrator := migrations.NewMigrator(db)
	defer log.Printf("Migration is compleat\n")

	if *flagMigrateUp {
		err = migrator.MigrateUp(&model.User{}, &model.Referral{}, &model.ReferralCode{})
	} else {
		err = migrator.MigrateDown(&model.User{}, &model.Referral{}, &model.ReferralCode{})
	}

	if err != nil {
		fmt.Println(err)
	}

}
