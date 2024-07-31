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
	ctx        = context.Background()
)

func seedHotel(name, location string) {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []*types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{

			Type:      types.DeluxeRoomType,
			BasePrice: 199.9,
		},
		{

			Type:      types.SeasideRoomType,
			BasePrice: 122.9,
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
	seedHotel("Bellucia", "France")
	seedHotel("The cozy hotel", "Netherlands")
	seedHotel("Don't die in your sleep", "London")
}

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)

	// Drop collections first to avoid duplicates when running seed
	err = hotelStore.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = roomStore.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
