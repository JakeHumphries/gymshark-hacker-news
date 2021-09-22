package main

import (
	"context"
	"log"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := consumer.ConnectDb(ctx)

	if err != nil {
		log.Fatal(errors.Wrap(err, "Main: "))
	}

	c := cron.New()

	c.AddFunc("0 30 * * * *", func() { consumer.Consume(ctx, mongoClient, consumer.HttpService{}) })

	c.Run()
}
