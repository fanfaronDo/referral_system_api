package main

import (
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"github.com/fanfaronDo/referral_system_api/pkg/app"
)

func main() {
	cnf := config.ConfigLoad()

	if err := app.Run(cnf); err != nil {
		fmt.Println(err)
	}
}
