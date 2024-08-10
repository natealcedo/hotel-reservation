package db

const (
	DBURI      = "mongodb://localhost:27017"
	DBNAME     = "hotel-reservations"
	TESTDBNAME = "test-hotel-reservations"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
