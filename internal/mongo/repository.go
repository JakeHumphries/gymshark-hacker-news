package mongo

import (
	"context"
	"fmt"

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
	log.Printf("Inserted Record: %v", item.Id)

	return &item, nil
}

// GetAllItems get all the items in the mongo database
func (r Repository) GetAllItems(ctx context.Context) {
	var items []models.Item
	database := r.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")
	cursor, err := itemsCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &items); err != nil {
		panic(err)
	}
	fmt.Println(items)
}
