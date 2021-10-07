package api

import (
	"context"
	"net/http"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// ItemReader is an interface for getting items
type ItemReader interface {
	GetAllItems(ctx context.Context) ([]models.Item, error)
	GetStories(ctx context.Context) ([]models.Item, error)
	GetJobs(ctx context.Context) ([]models.Item, error)
}

// Handler contains all the business logic for handling requests
type Handler struct {
	itemReader ItemReader
	ctx        context.Context
}

// New returns a new api handler
func New(itemReader ItemReader, ctx context.Context) *Handler {
	return &Handler{
		itemReader: itemReader,
		ctx:        ctx,
	}
}

// GetAllItems gets all the items in the item repository
func (h *Handler) GetAllItems(c echo.Context) (err error) {
	i, err := h.itemReader.GetAllItems(h.ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "handler"))
	}

	return c.JSON(http.StatusOK, i)
}

// GetStories gets all the items in the item repository with the type of story
func (h *Handler) GetStories(c echo.Context) (err error) {
	i, err := h.itemReader.GetStories(h.ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "handler"))
	}

	return c.JSON(http.StatusOK, i)
}

// GetJobs gets all the items in the item repository with the type of job
func (h *Handler) GetJobs(c echo.Context) (err error) {
	i, err := h.itemReader.GetJobs(h.ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "handler"))
	}

	return c.JSON(http.StatusOK, i)
}
