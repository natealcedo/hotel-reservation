package api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
	NumPersons int       `json:"numPersons"`
}

func (p *BookRoomParams) Validate() error {
	now := time.Now()
	if p.FromDate.Before(now) || p.TillDate.Before(now) {
		return fmt.Errorf("cannot book a room in the past")
	}
	return nil

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

	if err := params.Validate(); err != nil {
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

	ok, err = h.isRoomAvailableForBooking(c.Context(), roomId, params)

	if err != nil {
		return err
	}

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(genericResponse{
			Type: "error",
			Msg:  fmt.Sprintf("room %s is already booked", roomId.Hex()),
		})
	}

	insertedBooking, err := h.Booking.InsertBooking(c.Context(), &types.Booking{
		UserID:     user.ID,
		RoomID:     roomId,
		NumPersons: params.NumPersons,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(insertedBooking)
}

func (h *RoomHandler) isRoomAvailableForBooking(ctx context.Context, roomId primitive.ObjectID, params BookRoomParams) (bool, error) {
	where := bson.M{
		"roomID": roomId,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}

	bookings, err := h.Store.Booking.GetBookings(ctx, where)

	if err != nil {
		return false, err
	}

	ok := len(bookings) == 0
	return ok, nil
}

func (h *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms, err := h.Store.Room.GetRooms(c.Context(), nil)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}
