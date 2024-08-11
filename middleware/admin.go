package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/types"
)

func AdminAuth(ctx *fiber.Ctx) error {
	user, ok := ctx.Context().UserValue("user").(*types.User)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if !user.IsAdmin {
		return fiber.ErrUnauthorized
	}

	return ctx.Next()
}
