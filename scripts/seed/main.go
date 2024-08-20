package main

import (
	"context"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		Hotel:   hotelStore,
		Room:    db.NewMongoRoomStore(client, hotelStore),
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
	}

	james := fixtures.AddUser(store, "bronny", "james", false)
	hotel := fixtures.AddHotel(store, "Bellucia", "France", 5, nil)
	room := fixtures.AddRoom(store, "small", true, 89.99, hotel.ID)
	fixtures.AddBooking(store, james.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 3))
}
