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
	DB_NAME_TEST string
	SMTP_HOST   string
	SMTP_PORT   string
	SENDER_NAME string
	AUTH_EMAIL  string
	AUTH_PASSWORD  string
)

func Init() {
	godotenv.Load(".env")

	JWT_KEY = os.Getenv("JWT_KEY")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_NAME_TEST = os.Getenv("DB_NAME_TEST")
	SMTP_HOST   = os.Getenv("SMTP_HOST")
	SMTP_PORT   = os.Getenv("SMTP_PORT")
	SENDER_NAME = os.Getenv("SENDER_NAME")
	AUTH_EMAIL  = os.Getenv("AUTH_EMAIL")
	AUTH_PASSWORD  = os.Getenv("AUTH_PASSWORD")
}
