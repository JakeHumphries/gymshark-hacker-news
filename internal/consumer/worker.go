package consumer

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Worker is responsible for doing the work to save items
type Worker struct {
	itemProvider ItemProvider
	itemRepository ItemRepository
}

// NewWorker creates a new worker
func NewWorker(itemProvider ItemProvider, itemRepository ItemRepository) *Worker {
	return &Worker{
		itemProvider: itemProvider,
		itemRepository: itemRepository,
	}
}

func (w *Worker) run(ctx context.Context, idChan chan int, wg sync.WaitGroup) {
	for id := range idChan {
		item, err := w.itemProvider.GetItem(id)
		if err != nil {
			log.Print(errors.Wrap(err, "worker: "))
		} else if !item.Dead && !item.Deleted {
			_, err := w.itemRepository.SaveItem(ctx, *item)
			if err != nil {
				log.Print(errors.Wrap(err, "worker: "))
			}
		}
	}
	wg.Done()
}
