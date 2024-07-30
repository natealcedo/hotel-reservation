package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil || errors.Is(err, primitive.ErrInvalidHex) {
		return c.JSON(fiber.Map{
			"error": "not found",
		})
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return c.JSON([]any{})
	}

	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params *types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	errs := params.Validate()

	if len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	newUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(newUser)
}

func (h *UserHandler) HandleDeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	err := h.userStore.DeleteUserById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"deleted": id,
	})
}

func (h *UserHandler) HandlePutUserById(c *fiber.Ctx) error {
	var (
		params *types.UpdateUserParams
		userID = c.Params("id")
	)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err = c.BodyParser(&params); err != nil {
		return err
	}

	if errs := params.Validate(); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUserById(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(fiber.Map{"updated": userID})
}
