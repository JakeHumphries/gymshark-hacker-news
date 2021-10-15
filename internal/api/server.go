package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
)

type Server struct{
	Client protobufs.HackerNewsClient
}

func (s Server) Run(cfg *models.Config) {
	router := echo.New()
	router.HideBanner = true

	a := New(s.Client)

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
