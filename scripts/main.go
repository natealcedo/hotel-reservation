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

func seedUser(first, last, email string, isAdmin bool) {
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

}

func seedHotel(name, location string, rating int) {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []*types.Room{
		{
			Size:  "small",
			Price: 99.9,
		},
		{

			Size:  "normal",
			Price: 199.9,
		},
		{

			Size:  "kingsize",
			Price: 122.9,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedHotel("Bellucia", "France", 5)
	seedHotel("The cozy hotel", "Netherlands", 4)
	seedHotel("Don't die in your sleep", "London", 3)
	seedUser("Nate", "Alcedo", "natealcedo@gmail.com", false)
	seedUser("Lebron", "James", "lebron@gmail.com", false)
	seedUser("Bronny", "James", "bronny@gmail.com", true)
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
