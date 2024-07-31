package db

import (
	"context"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelColl = "hotels"

type HotelStore interface {
	Dropper

	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbName string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbName).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := s.coll.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
