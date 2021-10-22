package consumer

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/publisher"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Writer is an interface for saving items
type Writer interface {
	SaveItem(ctx context.Context, item models.Item) (*models.Item, error)
}

// Worker holds the execution logic that gets and saves items
type Worker struct {
	Provider publisher.Provider
	Writer   Writer
}

func (w Worker) Run(ctx context.Context, idChan chan int) {
	for id := range idChan {
		item, err := w.Provider.GetItem(id)
		if err != nil {
			log.Print(errors.Wrap(err, "worker"))
		} else if !item.Dead && !item.Deleted {
			_, err := w.Writer.SaveItem(ctx, *item)
			if err != nil {
				log.Print(errors.Wrap(err, "worker"))
			}
		}
	}
}
