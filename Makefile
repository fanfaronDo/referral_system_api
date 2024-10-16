
all: run

run:
	go run cmd/app/main.go

migrateup:
	go run cmd/migrate/main.go --up

migratedown:
	go run cmd/migrate/main.go --down