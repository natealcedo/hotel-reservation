package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	DBURI  = "mongodb://localhost:27017"
	DBNAME = "hotel-reservations"
)

func ToObjectId(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	return oid
}
