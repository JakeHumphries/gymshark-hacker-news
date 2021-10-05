package api

import (
	"context"
	"net/http"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Handler contains all the business logic for handling requests
type Handler struct {
	itemRepository consumer.ItemRepository
	ctx context.Context
}

// New returns a new api handler
func New(itemRepository consumer.ItemRepository) *Handler {
	return &Handler{
		itemRepository: itemRepository,
	}
}

// GetAll gets all the items in the item repository
func (h *Handler) GetAll(c echo.Context) (err error) {
	i, err := h.itemRepository.GetAllItems(h.ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.Wrap(err, "get all"))
	}

	return c.JSON(http.StatusOK, i)
}
