package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"google.golang.org/grpc"
)

type Server struct {
	Handler protobufs.HackerNewsServer
}

func (s Server) Run(cfg *models.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	protobufs.RegisterHackerNewsServer(srv, s.Handler)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}