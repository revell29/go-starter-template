package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	*AppConfig
	DBSystemConfig *DBConfig
	TenantConfig   *DBConfig
}

type AppConfig struct {
	NodePort        string
	NodeEnvironment string
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func GetConfig() *Config {
	godotenv.Load()

	config := &Config{
		AppConfig: &AppConfig{
			os.Getenv("NODE_PORT"),
			os.Getenv("NODE_ENV"),
		},
		DBSystemConfig: &DBConfig{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		},
		TenantConfig: &DBConfig{
			os.Getenv("DB_TENANT_HOST"),
			os.Getenv("DB_TENANT_PORT"),
			os.Getenv("DB_TENANT_USER"),
			os.Getenv("DB_TENANT_PASSWORD"),
			os.Getenv("DB_TENANT_NAME"),
		},
	}
	return config
}
