package db

import (
	"os"
)

var (
	DBURI  = "mongodb://localhost:27017"
	DBNAME = "hotel-reservations"
)

func init() {
	env := os.Getenv("ENV")
	if env == "test" {
		DBNAME = "test-hotel-reservations"
	}
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
