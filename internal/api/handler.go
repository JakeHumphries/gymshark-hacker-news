package api

import (
	"context"
	"io"

	grpc "github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/mappers"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Stream is a generic interface for grpc client streams
type Stream interface {
	Recv() (*protobufs.Item, error)
}

func readStream(s Stream) ([]models.Item, error) {
	var items []models.Item
	for {
		item, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		items = append(items, grpc.ToModel(item))
	}
	return items, nil
}

// GrpcHandler contains the logic that calls the grpc server
type GrpcHandler struct {
	client protobufs.HackerNewsClient
}

// New returns a new api router
func NewHandler(client protobufs.HackerNewsClient) *GrpcHandler {
	return &GrpcHandler{
		client: client,
	}
}

// GetAllItems gets all items in the grpc all stream
func (h GrpcHandler) GetAllItems(c context.Context) ([]models.Item, error) {
	s, err := h.client.All(c, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "calling grpc all")
	}

	items, err := readStream(s)
	if err != nil {
		return nil, errors.Wrap(err, "reading grpc all stream")
	}

	return items, nil
}

// GetStories gets all items in the grpc stories stream
func (h GrpcHandler) GetStories(c context.Context) ([]models.Item, error) {
	s, err := h.client.Stories(c, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "calling grpc stories")
	}

	items, err := readStream(s)
	if err != nil {
		return nil, errors.Wrap(err, "reading grpc stories stream")
	}

	return items, nil
}

// GetJobs gets all items in the grpc jobs stream
func (h GrpcHandler) GetJobs(c context.Context) ([]models.Item, error) {
	s, err := h.client.Jobs(c, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "calling grpc jobs")
	}

	items, err := readStream(s)
	if err != nil {
		return nil, errors.Wrap(err, "reading grpc jobs stream")
	}

	return items, nil
}
