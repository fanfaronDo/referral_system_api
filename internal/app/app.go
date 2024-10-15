package app

import (
	"context"
	"fmt"
	"github.com/fanfaronDo/referral_system_api/config"
	"github.com/fanfaronDo/referral_system_api/internal/entry"
	"github.com/fanfaronDo/referral_system_api/internal/handler"
	"github.com/fanfaronDo/referral_system_api/internal/service"
	"github.com/fanfaronDo/referral_system_api/internal/storage"
	"github.com/fanfaronDo/referral_system_api/migrations"
	"github.com/fanfaronDo/referral_system_api/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cnf *config.Config) error {
	ctx := context.Background()
	db, err := storage.NewPostgres(cnf.Postgres.Host, cnf)
	if err != nil {
		return err
	}
	migrator := migrations.NewMigrator(db)
	err = migrator.MigrateUp(&entry.User{}, &entry.Referral{})

	if err != nil {
		fmt.Println(err)
	}

	storage := storage.NewStorage(db)
	service := service.NewService(storage)
	handler := handler.NewHandler(service)
	routes := handler.InitRoutes()
	server := server.NewServer(cnf.HttpServer.Address+":"+cnf.HttpServer.Port, cnf, routes)

	go func() {
		if err = server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	defer server.Stop(ctx)

	log.Printf("Server started on %s\n", "http://"+cnf.HttpServer.Address+":"+cnf.HttpServer.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Shutting down server...\n")

	return nil
}
