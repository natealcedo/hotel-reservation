package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	RoomID     primitive.ObjectID `json:"roomID,omitempty" bson:"roomID,omitempty"`
	NumPersons int                `json:"numPersons,omitempty" bson:"numPersons,omitempty"`
	FromDate   time.Time          `json:"fromDate,omitempty" bson:"fromDate,omitempty"`
	TillDate   time.Time          `json:"tillDate,omitempty" bson:"tillDate,omitempty"`
	Canceled   bool               `json:"canceled" bson:"canceled"`
}
