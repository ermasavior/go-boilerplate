package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envPath = ".env"

var cfg Config

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

func Load() {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalln("error loading env variables", err)
	}

	cfg = Config{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func Get() Config {
	return cfg
}
