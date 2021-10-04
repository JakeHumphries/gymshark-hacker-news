package main

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/hackernews"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/mongo"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("loading .env file: %s", err)
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}

	mongoClient, err := mongo.ConnectDb(ctx, *cfg)
	if err != nil {
		log.Fatalf("connecting to db: %s", err)
	}

	c := cron.New()

	execute := func() {
		consumer.Execute(ctx, *cfg, mongo.Repository{Client: mongoClient}, hackernews.Api{})
	}

	execute()

	c.AddFunc(cfg.Cron, execute)

	c.Run()
}
