package internal

import (
	"cmp"
	"os"
)

type Config struct {
	Host  string
	Port  string
	DBDSN string
}

func ReadConfig() *Config {
	var cfg Config

	cfg.Host = cmp.Or(os.Getenv("RENTAL_SERVICE_HOST"), "0.0.0.0")
	cfg.Port = cmp.Or(os.Getenv("RENTAL_SERVICE_PORT"), "8080")
	cfg.DBDSN = cmp.Or(os.Getenv("RENTAL_SERVICE_DB_DSN"), "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable")

	return &cfg
}
