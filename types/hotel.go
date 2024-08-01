package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   int                  `json:"rating" bson:"rating"`
}

type RoomSize string

const (
	Small  RoomSize = "small"
	Normal          = "normal"
	King            = "king"
)

type Room struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Size    RoomSize           `json:"size" bson:"size"`
	Seaside bool               `json:"seaside" bson:"seaside"`
	Price   float64            `json:"price" bson:"price"`
	HotelID primitive.ObjectID `json:"hotelID" bson:"hotelID"`
}
