package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type BookingHandler struct {
	store *db.Store
}

type GetBookingsQueryParams struct {
	FromDate   time.Time `json:"fromDate,omitempty"`
	TillDate   time.Time `json:"tillDate,omitempty"`
	NumPersons int       `json:"numPersons,omitempty"`
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

// TODO: This needs to be admin authorized
func (h *BookingHandler) HandleGetBookings(ctx *fiber.Ctx) error {
	var params GetBookingsQueryParams
	if err := ctx.QueryParser(&params); err != nil {
		return err
	}

	fmt.Println(params.FromDate)

	filter := bson.M{}

	if params.NumPersons != 0 {
		filter["numPersons"] = bson.M{
			"$eq": params.NumPersons,
		}
	}

	if !params.FromDate.IsZero() {
		filter["fromDate"] = bson.M{
			"$gte": params.FromDate,
		}
	}

	if !params.TillDate.IsZero() {
		filter["tillDate"] = bson.M{
			"$lte": params.TillDate,
		}
	}

	bookings, err := h.store.Booking.GetBookings(ctx.Context(), filter)

	if err != nil {
		return err
	}

	return ctx.JSON(bookings)
}

// TODO: This needs to be user authorized - only the user who made the booking can see it
func (h *BookingHandler) HandleGetBookingByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	booking, err := h.store.Booking.GetBookingByID(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.JSON(booking)
}
