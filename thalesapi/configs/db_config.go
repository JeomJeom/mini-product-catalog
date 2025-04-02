package configs

import (
	"os"
)

type (
	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
)

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
