package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/api"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The address of the server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users/:id", userHandler.HandleGetUserById)
	apiV1.Get("/users", userHandler.HandleGetUsers)

	err = app.Listen(*listenAddr)

	if err != nil {
		panic(err)
	}
}
