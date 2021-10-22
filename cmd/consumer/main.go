package main

import (
	"context"
	"sync"
	"time"

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

	repo, err := mongo.NewRepository(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating mongo repository: %s", err)
	}

	q, err := queue.New(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating rabbitmq queue: %s", err)
	}
	defer q.Close()

	var wg sync.WaitGroup

	idChan := make(chan int)

	w := consumer.Worker{Provider: hackernews.Api{}, Writer: repo}

	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			w.Run(ctx, idChan)
		}()

	}

	q.Consume(idChan)

	close(idChan)

	wg.Wait()
}
