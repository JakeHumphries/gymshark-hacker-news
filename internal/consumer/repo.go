package consumer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDb(ctx context.Context) (*mongo.Client, error) {
	url := "mongodb://admin:admin@localhost:27017"

	fmt.Printf("connecting to mongodb at: %v", url)

	opts := options.Client()

	opts.ApplyURI(url)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "Connect DB: ")
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, errors.Wrap(err, "Connect DB: ")
	}

	fmt.Println("connected successfully to mongodb")

	return client, nil
}

func saveItem(ctx context.Context, mongoClient *mongo.Client, item Item) error {
	database := mongoClient.Database("hacker-news")
	itemsCollection := database.Collection("items")

	_, err := itemsCollection.InsertOne(ctx, item)
	if err != nil {
		return errors.Wrap(err, "Save item: ")
	}
	fmt.Printf("Inserted Record: %v \n", item.Id)

	return nil
}
