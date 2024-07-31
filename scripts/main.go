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

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)

	hotel := &types.Hotel{
		Name:     "Hotel California",
		Location: "California",
		Rooms:    []primitive.ObjectID{},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, hotel)

	if err != nil {
		log.Fatal(err)
	}

	rooms := []*types.Room{
		{
			HotelID:   insertedHotel.ID,
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{

			HotelID:   insertedHotel.ID,
			Type:      types.DeluxeRoomType,
			BasePrice: 199.9,
		},
		{

			HotelID:   insertedHotel.ID,
			Type:      types.SeasideRoomType,
			BasePrice: 122.9,
		},
	}

	var insertedRooms []*types.Room

	for _, room := range rooms {
		room, err := roomStore.InsertRoom(ctx, room)
		if err != nil {
			log.Fatal(err)
		}
		insertedRooms = append(insertedRooms, room)

	}

}
