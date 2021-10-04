package consumer

import (
	"context"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ItemProvider is an interface for getting hackernews data
type ItemProvider interface {
	GetTopStories() ([]int, error)
	GetItem(id int) (*models.Item, error)
}

// ItemRepository is an interface for saving items to persistance
type ItemRepository interface {
	SaveItem(ctx context.Context, item models.Item) (*models.Item, error)
	GetAllItems(ctx context.Context)
}

// Execute is the entry point for consumer service
func Execute(ctx context.Context, cfg models.Config, itemRepository ItemRepository, itemProvider ItemProvider) {
	ids, err := itemProvider.GetTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "execute"))
	}

	idChan := make(chan int)

	var wg sync.WaitGroup

	w := NewWorker(itemProvider, itemRepository)

	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			w.run(ctx, idChan)
		}()

	}

	for _, id := range ids {
		idChan <- id
	}
	close(idChan)

	wg.Wait()
}
