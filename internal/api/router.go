package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
)

// Handler contains all the business logic for handling requests
type Handler interface {
	GetAllItems(c context.Context) ([]models.Item, error)
	GetStories(c context.Context) ([]models.Item, error)
	GetJobs(c context.Context) ([]models.Item, error)
}


// Router contains all the routes for handling requests
type Router struct {
	handler Handler
}

// New returns a new api router
func NewRouter(handler Handler) *Router {
	return &Router{
		handler: handler,
	}
}

// All calls the GetAllItems route handler
func (r *Router) All(c echo.Context) (err error) {
	items, err := r.handler.GetAllItems(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling getAllitems: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}

// Stories calls the GetStories router handler
func (r *Router) Stories(c echo.Context) (err error) {
	items, err := r.handler.GetStories(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling getStories: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}

// Jobs calls the GetJobs router handler
func (r *Router) Jobs(c echo.Context) (err error) {
	items, err := r.handler.GetJobs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("calling getJobs: %s", err))
	}

	return c.JSON(http.StatusOK, items)
}
