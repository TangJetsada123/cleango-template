package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI    string
	jwtSecret   string
	groupID     string
	broadcastIP string
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	mongoURI = os.Getenv("MONGOURI")
	jwtSecret = os.Getenv("JWT_SECRET")
	groupID = os.Getenv("GROUPID")
	broadcastIP = os.Getenv("BROADCAST_IP")
}

func EnvMongoURI() string {
	return mongoURI
}

func JwtSecret() string {
	return jwtSecret
}

func GroupID() string {
	return groupID
}

func BroadcastIP() string {
	return broadcastIP
}
