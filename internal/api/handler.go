package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"

	grpc "github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/mappers"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
)

// ItemReader is an interface for getting items
type ItemReader interface {
	GetAllItems(ctx context.Context) ([]models.Item, error)
	GetStories(ctx context.Context) ([]models.Item, error)
	GetJobs(ctx context.Context) ([]models.Item, error)
}

// Stream is a generic interface for grpc client streams
type Stream interface {
	Recv() (*protobufs.Item, error)
}

// Handler contains all the business logic for handling requests
type Handler struct {
	client protobufs.HackerNewsClient
}

// New returns a new api handler
func New(client protobufs.HackerNewsClient) *Handler {
	return &Handler{
		client: client,
	}
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

// GetAllItems gets all items in the grpc all stream
func (h *Handler) GetAllItems(c echo.Context) (err error) {
	s, err := h.client.All(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling grpc all: %s", err))
	}

	items, err := readStream(s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("reading grpc all stream: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetStories gets all items in the grpc stories stream
func (h *Handler) GetStories(c echo.Context) (err error) {
	s, err := h.client.Stories(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling grpc stories: %s", err))
	}

	items, err := readStream(s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("reading grpc stories stream: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetJobs gets all items in the grpc jobs stream
func (h *Handler) GetJobs(c echo.Context) (err error) {
	s, err := h.client.Jobs(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling grpc jobs: %s", err))
	}

	items, err := readStream(s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("reading grpc jobs stream: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}
