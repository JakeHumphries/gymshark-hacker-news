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
}

// Execute is the entry point for consumer service
func Execute(ctx context.Context, itemRepository ItemRepository, itemProvider ItemProvider) {
	ids, err := itemProvider.GetTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "execute: "))
	}

	const workerCount = 10

	idChan := make(chan int)

	go populateIdChan(idChan, ids)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go worker(ctx, idChan, itemProvider, itemRepository, wg)
	}
	wg.Wait()
}

func populateIdChan(c chan int, ids []int) {
	for _, id := range ids {
		c <- id
	}
	close(c)
}
