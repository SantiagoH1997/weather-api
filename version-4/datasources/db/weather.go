package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	weatherCollection = "weather"
	indexKey          = "location_name"
)

// Open opens a connection to the DB
// It returns a MongoDB client, a Disconnect function, and an error
func Open(mongoURI, dbName string) (*mongo.Database, func(ctx context.Context) error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}
	db := client.Database(dbName)
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: indexKey, Value: bsonx.String("text")}}
	index.Options = options.Index().SetUnique(true)
	_, err = db.Collection(weatherCollection).Indexes().CreateOne(ctx, index)
	if err != nil {
		return nil, nil, err
	}
	return db, client.Disconnect, nil
}
