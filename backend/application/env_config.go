package application

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var dbHost string
var dbPort string
var dbUser string
var dbPassword string
var dbName string
var dbSSLMode string

func (a *App) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbSSLMode = os.Getenv("DB_SSLMODE")
}
