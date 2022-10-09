package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiKey    = ""
	SecretKey = ""
	Coin      = "BTCUSDT"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ApiKey = os.Getenv("API_KEY")
	SecretKey = os.Getenv("SECRET_KEY")
}
