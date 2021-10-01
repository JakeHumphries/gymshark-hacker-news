package mongo

import (
	"context"
	"fmt"
	"net/url"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDb tests the connection to mongo and returns a mongo client
func ConnectDb(ctx context.Context, cfg models.Config) (*mongo.Client, error) {

	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s", url.QueryEscape(cfg.DatabaseUser), url.QueryEscape(cfg.DatabasePassword), url.QueryEscape(cfg.DatabaseName), cfg.DatabasePort)

	log.Printf("connecting to mongodb at: %v \n", mongoUrl)

	opts := options.Client()

	opts.ApplyURI(mongoUrl)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect db")
	}

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("mongo refused to connect: %v %w", ctx.Err(), err)
		default:
			err := client.Ping(ctx, readpref.Primary())
			if err == nil {
				log.Print("mongo is now connected")
				return client, nil
			}
		}
	}
}
