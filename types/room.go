package types

import "go.mongodb.org/mongo-driver/bson/primitive"

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
