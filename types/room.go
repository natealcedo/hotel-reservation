package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomSize string

type Room struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// small, normal, kingsize
	Size      string             `json:"size" bson:"size"`
	Seaside   bool               `json:"seaside" bson:"seaside"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotelID" bson:"hotelID"`
	Available bool               `json:"-" bson:"available"`
}
