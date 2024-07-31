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

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The address of the server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	hotelStore := db.NewMongoHotelStore(client)
	hotelHandler := api.NewHotelHandler(db.NewMongoHotelStore(client), db.NewMongoRoomStore(client, hotelStore))

	app := fiber.New(config)
	apiV1 := app.Group("/api/v1")

	// Users
	apiV1.Get("/users/:id", userHandler.HandleGetUserById)
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUserById)
	apiV1.Put("/users/:id", userHandler.HandlePutUserById)

	// Hotels
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)

	err = app.Listen(*listenAddr)

	if err != nil {
		panic(err)
	}
}
