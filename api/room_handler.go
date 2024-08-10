package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
	NumPersons int       `json:"numPersons"`
}

type RoomHandler struct {
	*db.Store
}

func NewRoomHandler(s *db.Store) *RoomHandler {
	return &RoomHandler{
		Store: s,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(genericResponse{
			Msg:  "internal server error",
			Type: "error",
		})
	}

	booking := types.Booking{
		UserID:     user.ID,
		RoomID:     roomId,
		NumPersons: params.NumPersons,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
	}

	fmt.Printf("%+v\n", booking)

	return nil
}
