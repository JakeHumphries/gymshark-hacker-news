package main

import (
	"context"
	"log"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, err := consumer.ConnectDb(ctx)

	if err != nil {
		log.Fatal(errors.Wrap(err, "Main: "))
	}

	c := cron.New()

	consumer.Execute(consumer.MongoRepository{Client: mongoClient, Ctx: ctx}, consumer.HttpService{})

	c.AddFunc("0 30 * * * *", func() {
		consumer.Execute(consumer.MongoRepository{Client: mongoClient, Ctx: ctx}, consumer.HttpService{})
	})

	c.Run()
}
