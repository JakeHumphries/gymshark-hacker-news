package consumer

import (
	"log"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/pkg/models"
	"github.com/pkg/errors"
)

// DataGetter is an interface for getting hackernews data
type DataGetter interface {
	GetTopStories() ([]int, error)
	GetItem(id int) (*models.Item, error)
}

// ItemSaver is an interface for saving items to persistance
type ItemSaver interface {
	SaveItem(item models.Item) (*models.Item, error)
}

// Execute is the entry point for consumer service
func Execute(itemSaver ItemSaver, dataGetter DataGetter) {
	ids, err := dataGetter.GetTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "consume: "))
	}

	const workerCount = 10

	idChan := make(chan int)
	itemChan := make(chan models.Item)

	go populateIdChan(idChan, ids)

	go fanOutIds(workerCount, idChan, itemChan, dataGetter)

	for i := range itemChan {
		_, err := itemSaver.SaveItem(i)
		if err != nil {
			log.Print(errors.Wrap(err, "consume: "))
		}
	}
}

func fanOutIds(workerCount int, idChan chan int, itemChan chan models.Item, dataGetter DataGetter) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			for id := range idChan {
				func(id2 int) {
					item, err := dataGetter.GetItem(id2)
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
