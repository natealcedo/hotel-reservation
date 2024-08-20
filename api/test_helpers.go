package api

import (
	"context"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const (
	DBURI  = "mongodb://localhost:27017"
	DBNAME = "test-hotel-reservations"
)

type testDatabase struct {
	client *mongo.Client
	*db.Store
}

func setup(t *testing.T) *testDatabase {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBURI))

	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	hotelStore := db.NewMongoHotelStore(client)
	return &testDatabase{
		client: client,
		Store: &db.Store{
			User:    db.NewMongoUserStore(client),
			Booking: db.NewMongoBookingStore(client),
			Hotel:   db.NewMongoHotelStore(client),
			Room:    db.NewMongoRoomStore(client, hotelStore),
		},
	}

}

func (testingDb *testDatabase) tearDown(t *testing.T) {
	if err := testingDb.client.Database(db.DBNAME).Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}
