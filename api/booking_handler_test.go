package api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db/fixtures"
	"github.com/natealcedo/hotel-reservation/middleware"
	"github.com/natealcedo/hotel-reservation/types"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	user := fixtures.AddUser(db.Store, "admin", "admin", true)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "A", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 100.10, hotel.ID)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 1))
	bookingHandler := NewBookingHandler(db.Store)
	app := fiber.New()
	admin := app.Group("/", middleware.JWTAuthentication(db.Store.User), middleware.AdminAuth)
	admin.Get("/", bookingHandler.HandleGetBookings)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Fatal(err)
	}

	var bookings []*types.Booking

	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}

	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking but got %d", len(bookings))
	}

	if booking.ID != bookings[0].ID {
		t.Fatalf("expected %v but got %v", booking.ID, bookings[0].ID)
	}

	if booking.UserID != bookings[0].UserID {
		t.Fatalf("expected %v but got %v", booking.UserID, bookings[0].UserID)
	}
}

func TestNonAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	user := fixtures.AddUser(db.Store, "admin", "admin", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "A", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 100.10, hotel.ID)
	fixtures.AddBooking(db.Store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 1))
	bookingHandler := NewBookingHandler(db.Store)
	app := fiber.New()
	admin := app.Group("/", middleware.JWTAuthentication(db.Store.User), middleware.AdminAuth)
	admin.Get("/", bookingHandler.HandleGetBookings)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)

	if resp.StatusCode == fiber.StatusOK {
		t.Fatalf("expected %d but got %d", fiber.StatusInternalServerError, resp.StatusCode)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	user := fixtures.AddUser(db.Store, "admin", "admin", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "A", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 100.10, hotel.ID)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 1))
	bookingHandler := NewBookingHandler(db.Store)
	app := fiber.New()
	admin := app.Group("/", middleware.JWTAuthentication(db.Store.User))
	admin.Get("/:id", bookingHandler.HandleGetBookingByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	var bookingFromResponse types.Booking

	if err := json.NewDecoder(resp.Body).Decode(&bookingFromResponse); err != nil {
		t.Fatal(err)
	}

	if booking.ID != bookingFromResponse.ID {
		t.Fatalf("expected %v but got %v", booking.ID, bookingFromResponse.ID)
	}

	if booking.UserID != bookingFromResponse.UserID {
		t.Fatalf("expected %v but got %v", booking.UserID, bookingFromResponse.UserID)
	}
}

func TestNonAuthUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	user := fixtures.AddUser(db.Store, "admin", "admin", false)
	nonAuthUser := fixtures.AddUser(db.Store, "nonAuthUser", "nonAuthUser", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "A", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 100.10, hotel.ID)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 1))
	bookingHandler := NewBookingHandler(db.Store)
	app := fiber.New()
	admin := app.Group("/", middleware.JWTAuthentication(db.Store.User))
	admin.Get("/:id", bookingHandler.HandleGetBookingByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(nonAuthUser))
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected %d but got %d", fiber.StatusOK, resp.StatusCode)
	}
}
