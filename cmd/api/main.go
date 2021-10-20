package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/api"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("loading .env file: %s", err)
	}
}

func main() {
	cfg, err := models.GetConfig()
	if err != nil {
		log.Fatalf("loading config: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", cfg.GrpcHost, cfg.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connecting to grpc server: %s", err)
	}
	defer conn.Close()

	c := protobufs.NewHackerNewsClient(conn)

	s := api.Server{Client: c}

	s.Run(cfg)

}
