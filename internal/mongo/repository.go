package mongo

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository is a struct that exposes the mongo client and context
type Repository struct {
	Client *mongo.Client
}

// SaveItem saves items to the mongo database
func (r Repository) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
	database := r.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")

	opts := options.Update().SetUpsert(true)

	update := bson.M{
		"$set": item,
	}
	_, err := itemsCollection.UpdateOne(ctx, bson.M{"id": item.Id}, update, opts)
	if err != nil {
		return nil, errors.Wrap(err, "save item: ")
	}
	log.Print("Inserted Record: %v \n", item.Id)

	return &item, nil
}
