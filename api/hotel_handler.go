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

type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var queryParams HotelQueryParams

	if err := c.QueryParser(&queryParams); err != nil {
		return err
	}

	hotels, err := h.hotelStore.GetHotels(c.Context(), nil)

	if err != nil {
		return err
	}

	if len(hotels) == 0 {
		return c.JSON([]any{})
	}

	return c.JSON(hotels)

}
