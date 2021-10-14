package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"google.golang.org/grpc"
)

type Server struct {
	Handler protobufs.HackerNewsServer
}

func (s Server) Run() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	protobufs.RegisterHackerNewsServer(srv, s.Handler)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}