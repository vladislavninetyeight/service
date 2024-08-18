package postgresql

import "os"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func GetConfig() *Config {
	return &Config{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Database: os.Getenv("DATABASE_NAME"),
	}
}
