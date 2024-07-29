package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The address of the server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users/:id", api.HandleGetUserById)
	apiV1.Get("/users", api.HandleGetUsers)

	err := app.Listen(*listenAddr)

	if err != nil {
		panic(err)
	}
}
