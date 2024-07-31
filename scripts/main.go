package main

import (
	"context"
	"fmt"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
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

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	hotel := &types.Hotel{
		Name:     "Hotel California",
		Location: "California",
	}
	//_ := &types.Room{
	//	Type:      types.SingleRoomType,
	//	BasePrice: 99.9,
	//}

	insertedHotel, err := hotelStore.InsertHotel(ctx, hotel)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedHotel)

}
