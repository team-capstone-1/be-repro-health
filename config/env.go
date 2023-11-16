package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_KEY     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
)

func init() {
	godotenv.Load(".env")

	JWT_KEY = os.Getenv("JWT_KEY")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
}
