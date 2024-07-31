package db

import (
	"context"
	"fmt"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	Dropper

	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(context.Context, bson.M, bson.M) error
	GetHotels(context.Context) ([]*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("hotels"),
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

func (s *MongoHotelStore) Drop(ctx context.Context) error {
	fmt.Printf("---- Dropping %s collection ----\n", s.coll.Name())
	return s.coll.Drop(ctx)
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoHotelStore) GetHotels(ctx context.Context) ([]*types.Hotel, error) {
	res, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel
	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}
