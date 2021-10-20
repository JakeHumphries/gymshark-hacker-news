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
	provider publisher.Provider
	writer   Writer
	idChan   chan int
}

func (w Worker) Run(ctx context.Context) {
	for id := range w.idChan {
		item, err := w.provider.GetItem(id)
		if err != nil {
			log.Print(errors.Wrap(err, "worker"))
		} else if !item.Dead && !item.Deleted {
			_, err := w.writer.SaveItem(ctx, *item)
			if err != nil {
				log.Print(errors.Wrap(err, "worker"))
			}
		}
	}
}
