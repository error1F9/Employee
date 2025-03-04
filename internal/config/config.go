package config

import "os"

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Port struct {
	Port string
}

type AppConfig struct {
	DB   DB
	Port Port
}

func NewConfig() AppConfig {
	return AppConfig{
		DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Port{os.Getenv("DB_PORT")},
	}
}
