package main

import (
	"context"
	"log"
	"time"

	"cleango-template/api/route"
	"cleango-template/internal/configs"
	"cleango-template/internal/handler"
	"cleango-template/internal/repository"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := configs.EnvMongoURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("your_db_name")
	userRepo := repository.NewUserRepository(db, "users")
	userHandler := handler.NewUserHandler(userRepo)

	e := echo.New()

	route.UserRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
