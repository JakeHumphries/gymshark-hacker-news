package grpc

import (
	"context"
	"fmt"

	grpc "github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/mappers"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ItemReader is an interface for getting items
type ItemReader interface {
	GetAllItems(ctx context.Context) ([]models.Item, error)
	GetStories(ctx context.Context) ([]models.Item, error)
	GetJobs(ctx context.Context) ([]models.Item, error)
}

// Handler contains all the business logic for handling requests
type Handler struct {
	protobufs.UnimplementedHackerNewsServer
	itemReader ItemReader
}

// New returns a new api handler
func New(itemReader ItemReader, ctx context.Context) *Handler {
	return &Handler{
		itemReader: itemReader,
	}
}

// All gets all the items in the item repository
func (h *Handler) All(e *emptypb.Empty, s protobufs.HackerNews_AllServer) error {
	items, err := h.itemReader.GetAllItems(s.Context())
	if err != nil {
		return fmt.Errorf("getAllItems: %w", err)
	}

	for _, item := range items {
		s.Send(grpc.ToProto(item))
	}

	return nil
}

// Stories gets all the items in the item repository with the type of story
func (h *Handler) Stories(e *emptypb.Empty, s protobufs.HackerNews_StoriesServer) error {
	items, err := h.itemReader.GetStories(s.Context())
	if err != nil {
		return fmt.Errorf("getStories: %w", err)
	}

	for _, item := range items {
		s.Send(grpc.ToProto(item))
	}

	return nil
}

// Jobs gets all the items in the item repository with the type of job
func (h *Handler) Jobs(e *emptypb.Empty, s protobufs.HackerNews_JobsServer) error {
	items, err := h.itemReader.GetJobs(s.Context())
	if err != nil {
		return fmt.Errorf("getJobs: %w", err)
	}

	for _, item := range items {
		s.Send(grpc.ToProto(item))
	}

	return nil
}
