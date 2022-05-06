package main

import (
	"fmt"
	"log"

	"github.com/artacone/go-url-short/internal/server"
	"github.com/caarlos0/env/v6"
)

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%q", err)
	}

	dsn := buildDSN(cfg)
	server.Run(dsn)
}

type config struct {
	DBUse      bool   `env:"DB_USE" envDefault:"true"`
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
}

func buildDSN(cfg config) string {
	if !cfg.DBUse {
		return ""
	}
	dsnTemplate := "host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(dsnTemplate, cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
}
