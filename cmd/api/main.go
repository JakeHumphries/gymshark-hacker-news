package main

import (
	"fmt"
	"net/http"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
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
	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}

	router := echo.New()
	router.HideBanner = true

	router.GET("/_healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	addr := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	if err := router.Start(addr); err != nil {
		log.Fatalf("starting server %s", err)
	}
}
