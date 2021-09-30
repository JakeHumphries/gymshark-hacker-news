package mongo

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	log "github.com/sirupsen/logrus"
)

// ConnectDb tests the connection to mongo and returns a mongo client
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

	log.Infof("connecting to mongodb at: %v \n", url)

	opts := options.Client()

	opts.ApplyURI(url)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect db: ")
	}

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("mongo refused to connect: %v %w", ctx.Err(), err)
		default:
			err := client.Ping(ctx, readpref.Primary())
			if err == nil {
				log.Infof("mongo is now connected")
				return client, nil
			}
		}
	}
}