package api

import (
	"context"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const testDBName = "test-hotel-reservation"

type testDb struct {
	db.UserStore
}

func setup(t *testing.T) *testDb {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))

	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	return &testDb{
		UserStore: db.NewMongoUserStore(client, testDBName),
	}

}

func (testingDb *testDb) tearDown(t *testing.T) {
	if err := testingDb.UserStore.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

}
