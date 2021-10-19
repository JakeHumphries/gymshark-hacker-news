package main

import (
	"context"

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

	q, err := queue.New(ctx)
	if err != nil {
		log.Fatalf("creating rabbitmq queue %s", err)
	}
	defer q.Close()

	c := cron.New()

	run := func() {
		publisher.Run(q, hackernews.Api{})
	}

	run()

	c.AddFunc(cfg.Cron, run)

	c.Run()
}
