package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"time"
)

type BookRoomParams struct {
	FromDate time.Time `json:"fromDate"`
	TillDate time.Time `json:"tillDate"`
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
	//roomOID, err := primitive.ObjectIDFromHex(c.Params("id"))
	//
	//if err != nil {
	//	return err
	//}

	return nil
}
