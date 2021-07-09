package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Database DatabaseConfig
	Jwt      JwtConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JwtConfig struct {
	SigningKey string
	Audience   string
	Issuer     string
}

func GetEnv() Config {
	return Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_USER_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		Jwt: JwtConfig{
			SigningKey: os.Getenv("JWT_SIGNING_KEY"),
			Audience:   os.Getenv("JWT_AUDIENCE"),
			Issuer:     os.Getenv("JWT_ISSUER"),
		},
	}
}
