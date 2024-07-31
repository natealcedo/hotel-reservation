package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
)

type HotelHandler struct {
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hs,
		roomStore:  rs,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.GetHotels(c.Context())

	if err != nil {
		return err
	}

	if len(hotels) == 0 {
		return c.JSON([]any{})
	}

	return c.JSON(hotels)

}
