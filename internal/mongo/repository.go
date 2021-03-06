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

// NewRepository creates a new repository
func NewRepository(ctx context.Context, cfg models.Config) (*Repository, error) {
	mongoClient, err := ConnectDb(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to db")
	}

	return &Repository{
		Client: mongoClient,
	}, nil
}

// SaveItem saves items to the mongo database
func (w Repository) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
	database := w.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")

	opts := options.Update().SetUpsert(true)

	update := bson.M{
		"$set": item,
	}
	_, err := itemsCollection.UpdateOne(ctx, bson.M{"id": item.Id}, update, opts)
	if err != nil {
		return nil, errors.Wrap(err, "repository: save item")
	}
	log.Printf("Inserted Record: %v", item.Id)

	return &item, nil
}

// GetAllItems get all the items in the mongo database
func (r Repository) GetAllItems(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	database := r.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")
	cursor, err := itemsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "repository: get all items")
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, errors.Wrap(err, "repository: get all items")
	}
	return items, nil
}

// GetStories get all the items in the mongo database with the type of story
func (r Repository) GetStories(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	database := r.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")
	cursor, err := itemsCollection.Find(ctx, bson.M{"type": "story"})
	if err != nil {
		return nil, errors.Wrap(err, "repository: get stories")
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, errors.Wrap(err, "repository: get stories")
	}
	return items, nil
}

// GetJobs get all the items in the mongo database with the type of job
func (r Repository) GetJobs(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	database := r.Client.Database("hacker-news")
	itemsCollection := database.Collection("items")
	cursor, err := itemsCollection.Find(ctx, bson.M{"type": "job"})
	if err != nil {
		return nil, errors.Wrap(err, "repository: get jobs")
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, errors.Wrap(err, "repository: get jobs")
	}
	return items, nil
}
