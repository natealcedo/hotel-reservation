package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
)

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	store *db.Store
}

func NewAuthHandler(store *db.Store) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var authParams *AuthParams
	if err := c.BodyParser(&authParams); err != nil {
		return err
	}
	fmt.Println(authParams)
	return nil
}
