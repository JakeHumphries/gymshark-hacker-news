package main

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/hackernews"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/mongo"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/queue"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("loading .env file %s", err)
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config %s", err)
	}

	repo, err := mongo.NewRepository(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating mongo repository %s", err)
	}

	q, err := queue.New(ctx)
	if err != nil {
		log.Fatalf("creating rabbitmq queue %s", err)
	}
	defer q.Close()

	consumer.Run(ctx, cfg, q, hackernews.Api{}, repo)
}
