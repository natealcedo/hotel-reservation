package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/api"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dbUri          = "mongodb://localhost:27017"
	dbName         = "hotel-reservations"
	userCollection = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbName).Collection(userCollection)
	user := &types.User{
		FirstName: "Nate",
		LastName:  "Alcedo",
	}

	_, err = coll.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err)
	}

	var usr types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&usr); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", usr)

	listenAddr := flag.String("listenAddr", ":3000", "The address of the server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users/:id", api.HandleGetUserById)
	apiV1.Get("/users", api.HandleGetUsers)

	err = app.Listen(*listenAddr)

	if err != nil {
		panic(err)
	}
}
