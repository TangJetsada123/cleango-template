package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return os.Getenv("MONGOURI")
}

func JwtSecret() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	return os.Getenv("JWT_SECRET")
}

func GroupID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	return os.Getenv("GroupID")
}

func BroadcastIP() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	return os.Getenv("BROADCAST_IP")
}
