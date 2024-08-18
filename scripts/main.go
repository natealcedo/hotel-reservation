package main

import (
	"context"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	roomStore  *db.MongoRoomStore
	hotelStore *db.MongoHotelStore
	userStore  *db.MongoUserStore
	ctx        = context.Background()
)

func seedUser(first, last, email string, isAdmin bool) *types.User {
	user, err := types.NewUserFromParams(&types.CreateUserParams{
		FirstName: first,
		LastName:  last,
		Email:     email,
		Password:  "supersecurepassword",
		IsAdmin:   isAdmin,
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.InsertUser(ctx, user)

	if err != nil {
		log.Fatal(err)
	}

	return user

}

func seedRoom(size string, seaside bool, price float64, roomID primitive.ObjectID) *types.Room {
	room := &types.Room{
		Size:    size,
		Price:   price,
		Seaside: seaside,
		HotelID: roomID,
	}

	_, err := roomStore.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return room
}

func seedHotel(name, location string, rating int) *types.Hotel {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func main() {
	seedUser("Nate", "Alcedo", "natealcedo@gmail.com", false)
	seedUser("Lebron", "James", "lebron@gmail.com", false)
	seedUser("Bronny", "James", "bronny@gmail.com", false)
	seedUser("Admin", "Admin", "admin@admin.com", true)
	seedHotel("Bellucia", "France", 5)
	seedHotel("The cozy hotel", "Netherlands", 4)
	hotel := seedHotel("Don't die in your sleep", "London", 3)
	seedRoom("small", true, 89.99, hotel.ID)
	seedRoom("medium", true, 189.99, hotel.ID)
	seedRoom("large", false, 289.99, hotel.ID)
}

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)

	// Drop collections first to avoid duplicates when running seed
	err = hotelStore.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = roomStore.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = userStore.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
