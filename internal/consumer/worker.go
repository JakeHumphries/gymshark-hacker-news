package consumer

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Worker is responsible for doing the work to save items
type Worker struct {
	itemProvider   ItemProvider
	itemWriter ItemWriter
}

// NewWorker creates a new worker
func NewWorker(itemProvider ItemProvider, itemWriter ItemWriter) *Worker {
	return &Worker{
		itemProvider:   itemProvider,
		itemWriter: itemWriter,
	}
}

func (w *Worker) run(ctx context.Context, idChan chan int) {
	for id := range idChan {
		item, err := w.itemProvider.GetItem(id)
		if err != nil {
			log.Print(errors.Wrap(err, "worker"))
		} else if !item.Dead && !item.Deleted {
			_, err := w.itemWriter.SaveItem(ctx, *item)
			if err != nil {
				log.Print(errors.Wrap(err, "worker"))
			}
		}
	}
}
