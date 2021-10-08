package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/api"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/mongo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
		log.Fatalf("loading config: %s", err)
	}

	repo, err := mongo.NewRepository(ctx, *cfg)
	if err != nil {
		log.Fatalf("creating mongo repository %s", err)
	}

	cacheReader := api.NewCacheReader(repo, *cfg)

	router := echo.New()
	router.HideBanner = true

	a := api.New(cacheReader, ctx)

	router.GET("/all", a.GetAllItems)

	router.GET("/stories", a.GetStories)

	router.GET("/jobs", a.GetJobs)

	router.GET("/_healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	addr := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	if err := router.Start(addr); err != nil {
		log.Fatalf("starting server %s", err)
	}
}
