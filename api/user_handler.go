package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/types"
)

func HandleGetUserById(c *fiber.Ctx) error {
	u := &types.User{
		ID:        "52fb62f9-4258-4624-857c-e2c51e0f3632",
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(u)
}

func HandleGetUsers(c *fiber.Ctx) error {
	users := &[]types.User{
		{
			ID:        "asdf123",
			FirstName: "Nate",
			LastName:  "Alcedo",
		},
		{
			ID:        "asdf123",
			FirstName: "Emmanuel",
			LastName:  "Alcedo",
		},
		{
			ID:        "asdf123",
			FirstName: "Evelyn",
			LastName:  "Fontentot",
		},
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
