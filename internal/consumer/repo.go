package consumer

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBRepository interface {
	Save(document interface{}) error
}

type MongoRepository struct {
	Client *mongo.Client
	Ctx context.Context
}

func (mr MongoRepository) Save(document interface{}) error {
	database := mr.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")

	_, err := itemsCollection.InsertOne(mr.Ctx, document)
	if err != nil {
		return errors.Wrap(err, "Save item: ")
	}

	return nil
}

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
