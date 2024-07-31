package db

import (
	"context"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomColl = "rooms"

type RoomStore interface {
	Dropper

	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbName string) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(dbName).Collection(roomColl),
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.coll.InsertOne(ctx, room)

	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)
	return room, nil
}
