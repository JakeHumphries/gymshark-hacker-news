package consumer

import (
	"log"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
)

// ItemProvider is an interface for getting hackernews data
type ItemProvider interface {
	GetTopStories() ([]int, error)
	GetItem(id int) (*models.Item, error)
}

// ItemRepository is an interface for saving items to persistance
type ItemRepository interface {
	SaveItem(item models.Item) (*models.Item, error)
}

// Execute is the entry point for consumer service
func Execute(itemRepository ItemRepository, itemProvider ItemProvider) {
	ids, err := itemProvider.GetTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "consume: "))
	}

	const workerCount = 10

	idChan := make(chan int)
	itemChan := make(chan models.Item)

	go populateIdChan(idChan, ids)

	go fanOutIds(workerCount, idChan, itemChan, itemProvider)

	for i := range itemChan {
		_, err := itemRepository.SaveItem(i)
		if err != nil {
			log.Print(errors.Wrap(err, "consume: "))
		}
	}
}

func fanOutIds(workerCount int, idChan chan int, itemChan chan models.Item, itemProvider ItemProvider) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			for id := range idChan {
				func(id2 int) {
					item, err := itemProvider.GetItem(id2)
					if err != nil {
						log.Print(errors.Wrap(err, "fan out ids: "))
					} else if !item.Dead && !item.Deleted {
						itemChan <- *item
					}
				}(id)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(itemChan)
}

func populateIdChan(c chan int, ids []int) {
	for _, id := range ids {
		c <- id
	}
	close(c)
}
