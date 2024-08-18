package main

import (
	"context"
	"fmt"
	"github.com/natealcedo/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Database(db.DBNAME).Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s database dropped\n", db.DBNAME)
}
