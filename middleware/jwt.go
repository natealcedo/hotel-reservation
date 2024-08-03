package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("Anauthorized")
	}

	fmt.Println(token)

	if err := c.Next(); err != nil {
		return err
	}

	return nil

}
