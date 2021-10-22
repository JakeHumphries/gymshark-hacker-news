package main

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/mongo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("loading .env file %s", err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}

	repo, err := mongo.NewRepository(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating mongo repository: %s", err)
	}

	cacheReader := grpc.NewCacheReader(repo, *cfg)

	router := echo.New()
	router.HideBanner = true

	h := grpc.New(cacheReader, ctx)

	s := grpc.Server{Handler: h}

	s.Run(cfg)
}
