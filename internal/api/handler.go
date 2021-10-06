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
}

// Handler contains all the business logic for handling requests
type Handler struct {
	itemReader ItemReader
	ctx context.Context
}

// New returns a new api handler
func New(itemReader ItemReader, ctx context.Context) *Handler {
	return &Handler{
		itemReader: itemReader,
		ctx: ctx,
	}
}

// GetAll gets all the items in the item repository
func (h *Handler) GetAll(c echo.Context) (err error) {
	i, err := h.itemReader.GetAllItems(h.ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "get all"))
	}

	return c.JSON(http.StatusOK, i)
}
