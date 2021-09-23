package main

import (
	"context"
	"log"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("loading .env file: %s", err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, err := consumer.ConnectDb(ctx)

	if err != nil {
		log.Fatalf("connecting to db: %s", err)
	}

	c := cron.New()

	execute := func() {
		consumer.Execute(consumer.MongoRepository{Client: mongoClient, Ctx: ctx}, consumer.HttpService{})
	}

	execute()

	c.AddFunc("0 30 * * * *", execute)

	c.Run()
}
