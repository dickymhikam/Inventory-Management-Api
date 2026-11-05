package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load")
	}

	return &Config{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_Port"),
		DBName:     os.Getenv("DB_Name"),
		DBUser:     os.Getenv("DB_User"),
		DBPassword: os.Getenv("DB_Password"),
	}, nil
}
