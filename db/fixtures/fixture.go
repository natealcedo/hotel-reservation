package fixtures

import (
	"context"
	"fmt"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddUser(store *db.Store, firstName, lastName string, admin bool) *types.User {
	user, err := types.NewUserFromParams(&types.CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     fmt.Sprintf("%s@%s.com", firstName, lastName),
		Password:  fmt.Sprintf("%s_%s", firstName, lastName),
		IsAdmin:   admin,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedUser, err := store.User.InsertUser(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddHotel(store *db.Store, name, location string, rating int, rooms []primitive.ObjectID) *types.Hotel {
	if rooms == nil {
		rooms = []primitive.ObjectID{}
	}
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rating:   rating,
		Rooms:    rooms,
	}

	insertedHotel, err := store.Hotel.InsertHotel(context.Background(), hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddRoom(store *db.Store, size string, seaside bool, price float64, hotelID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Price:   price,
		Seaside: seaside,
		HotelID: hotelID,
	}

	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, from, till time.Time) *types.Booking {
	booking := &types.Booking{
		UserID:     userID,
		RoomID:     roomID,
		NumPersons: 0,
		FromDate:   from,
		TillDate:   till,
	}

	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}
