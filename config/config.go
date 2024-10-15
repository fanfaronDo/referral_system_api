package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	HttpServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HttpServer struct {
	Address     string        `env:"HTTP_SERVER_ADDRESS" env-default:"127.0.0.1"`
	Port        string        `env:"HTTP_SERVER_PORT" env-default:"8080"`
	Timeout     time.Duration `env:"HTTP_SERVER_TIMEOUT" env-default:"3s"`
	IdleTimeout time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"30s"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Database string `env:"POSTGRES_DATABASE" env-default:"postgres"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

func ConfigLoad() *Config {
	var config Config
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n")
		os.Exit(1)
	}

	config.HttpServer.Address = getEnv("HTTP_SERVER_ADDRESS", "127.0.0.1")
	config.HttpServer.Port = getEnv("HTTP_SERVER_PORT", "8080")
	config.HttpServer.Timeout = getDurationEnv("HTTP_SERVER_TIMEOUT", 10*time.Second)
	config.HttpServer.IdleTimeout = getDurationEnv("HTTP_SERVER_IDLE_TIMEOUT", 30*time.Second)

	config.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	config.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	config.Postgres.User = getEnv("POSTGRES_USER", "postgres")
	config.Postgres.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	config.Postgres.Database = getEnv("POSTGRES_DATABASE", "postgres")
	config.Postgres.SSLMode = getEnv("POSTGRES_SSL_MODE", "disable")

	return &config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	value := getEnv(key, defaultValue.String())
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration for %s: %s, using default: %v\n", key, value, defaultValue)
		return defaultValue
	}
	return duration
}
