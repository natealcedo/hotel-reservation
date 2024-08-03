package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthParams struct {
	Email    string `json:"email" `
	Password string `json:"password" `
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
	// When using body parser, the variable has to be a struct, not a pointer to a struct
	// When passing the variable to c.BodyParser, it has to be a pointer to the variable
	// Not doing so will result in a nil dereference error
	var authParams AuthParams
	if err := c.BodyParser(&authParams); err != nil {
		return err
	}

	user, err := h.store.User.GetUserByEmail(c.Context(), authParams.Email)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}
		return err
	}

	return c.JSON(user)
}
