package consumer

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	url := "mongodb://admin:admin@localhost:27017"

	fmt.Printf("connect to mongodb at: %v", url)

	opts := options.Client()

	opts.ApplyURI(url)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		msg := fmt.Sprintf("invalid mongodb connection: %v", err)
		fmt.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Printf("connection test to mongodb client failed: %v", err)
		return nil, fmt.Errorf(fmt.Sprintf("connection test to mongodb client failed: %v", err))
	}

	fmt.Println("connected successfully to mongodb")

	return client, nil
}