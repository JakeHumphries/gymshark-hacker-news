package consumer

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbRepository interface {
	SaveItem(item Item) error
}

type MongoRepository struct {
	Client *mongo.Client
	Ctx    context.Context
}

func (mr MongoRepository) SaveItem(item Item) error {
	database := mr.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")

	_, err := itemsCollection.InsertOne(mr.Ctx, item)
	if err != nil {
		return errors.Wrap(err, "Save item: ")
	}
	fmt.Printf("Inserted Record: %v \n", item.Id)

	return nil
}

func ConnectDb(ctx context.Context) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Printf("connecting to mongodb at: %v \n", url)

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
