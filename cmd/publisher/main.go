package main

import (
	"context"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/hackernews"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/publisher"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/queue"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("loading .env file %s", err)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}

	q, err := queue.New(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating rabbitmq queue: %s", err)
	}
	defer q.Close()

	c := cron.New()

	run := func() {
		if err := publisher.Run(q, hackernews.Api{}); err != nil {
			log.Fatalf("running publisher: %s", err)
		}
	}

	run() // run is called once here on execution to initially publish the hn data

	c.AddFunc(cfg.Cron, run)

	c.Run()
}
