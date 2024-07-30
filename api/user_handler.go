package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
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
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params *types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	errors := params.Validate()

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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

//func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
//	id := c.Params("id")
//	err := h.userStore.DeleteUser(c.Context(), id)
//	if err != nil {
//		return fiber.NewError(fiber.StatusNotFound, err.Error())
//	}
//	return c.SendStatus(fiber.StatusNoContent)
//}
