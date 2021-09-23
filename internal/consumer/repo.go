package consumer

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseRepository interface {
	SaveItem(item Item) (*Item, error)
}

type MongoRepository struct {
	Client *mongo.Client
	Ctx    context.Context
}

func (mr MongoRepository) SaveItem(item Item) (*Item, error) {
	database := mr.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")

	opts := options.Update().SetUpsert(true)

	update := bson.M{
        "$set": item,
    }
	_, err := itemsCollection.UpdateOne(mr.Ctx, bson.M{"id": item.Id}, update, opts)

	if err != nil {
		return nil, errors.Wrap(err, "save item: ")
	}
	fmt.Printf("Inserted Record: %v \n", item.Id)

	return &item, nil
}

func ConnectDb(ctx context.Context) (*mongo.Client, error) {
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("err: env var database user doesnt exist")
	}
	pass, exists := os.LookupEnv("DB_PASS")
	if !exists {
		return nil, errors.New("err: env var database pass doesnt exist")
	}
	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, errors.New("err: env var database name doesnt exist")
	}
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("err: env var database port doesnt exist")
	}

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, name, port)

	fmt.Printf("connecting to mongodb at: %v \n", url)

	opts := options.Client()

	opts.ApplyURI(url)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect db: ")
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, errors.Wrap(err, "connect db: ")
	}

	fmt.Println("connected successfully to mongodb")

	return client, nil
}
