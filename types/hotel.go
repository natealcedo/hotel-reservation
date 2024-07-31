package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeasideRoomType
	DeluxeRoomType
)

type Room struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type      RoomType           `json:"type" bson:"type"`
	BasePrice float64            `json:"basePrice" bson:"basePrice"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotelID" bson:"hotelID"`
}
