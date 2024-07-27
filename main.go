package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/api"
	"net/http"
	"os"
)

func handleRoot(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hello, World!",
	})

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/users/:id", api.HandleGetUserById)
	apiV1.Get("/users", api.HandleGetUsers)

	err := app.Listen(":" + port)

	if err != nil {
		panic(err)
	}
}
