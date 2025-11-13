package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DbConfig struct {
	Dsn string
}

type LogConfig struct {
	Level string
}

type Config struct {
	App AppConfig
	Db  DbConfig
	Log LogConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "splitify"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Db: DbConfig{
			Dsn: getEnv("DB_DSN", "postgres://postgres:postgres@localhost:5432/splitify?sslmode=disable"),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return fallback
	}
}
