package main

import (
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"github.com/fanfaronDo/referral_system_api/pkg/model"
	"github.com/fanfaronDo/referral_system_api/pkg/storage"
	"log"
	"time"
)

var code = model.ReferralCode{
	Code:           "ASDDWAWDSawdawd",
	ExpirationTime: time.Second * 10,
	IsActive:       true,
	UserId:         10,
}

func main() {
	cnf := config.ConfigLoad()
	db, err := storage.NewPostgres(cnf.Postgres.Host, cnf)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}

	s := storage.NewReferralCode(db)
	c, e := s.GetReferralCode("3cbed43c")
	if e != nil {
		log.Fatalf("%s", e.Error())
	}

	s.UpdateReferralCodeStatus(c, false)
	fmt.Println(c)
}
