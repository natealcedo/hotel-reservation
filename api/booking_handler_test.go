package api

import (
	"fmt"
	"github.com/natealcedo/hotel-reservation/db/fixtures"
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

	fmt.Println(booking)
}
