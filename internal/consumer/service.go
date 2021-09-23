package consumer

import (
	"log"
	"sync"

	"github.com/pkg/errors"
)

type Item struct {
	By          string `bson:"by"`
	Title       string `bson:"title"`
	Url         string `bson:"url"`
	Text        string `bson:"text"`
	ItemType    string `bson:"itemType" json:"type"`
	Descendants int    `bson:"descendants"`
	Id          int    `bson:"id"`
	Score       int    `bson:"score"`
	Time        int    `bson:"time"`
	Parent      int    `bson:"parent"`
	Poll        int    `bson:"poll"`
	Kids        []int  `bson:"kids"`
	Parts       []int  `bson:"parts"`
	Deleted     bool   `bson:"deleted"`
	Dead        bool   `bson:"dead"`
}

// Execute - Entry point for consumer service
func Execute(dbRepo DbRepository, dataService DataService) {
	ids, err := dataService.getTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Consume: "))
	}

	const workerCount = 10

	idChan := make(chan int)
	itemChan := make(chan Item)

	go populateIdChan(idChan, ids)

	go fanOutIds(workerCount, idChan, itemChan, dataService)

	for i := range itemChan {
		err := dbRepo.SaveItem(i)
		if err != nil {
			log.Print(errors.Wrap(err, "Consume: "))
		}
	}
}

func fanOutIds(workerCount int, idChan chan int, itemChan chan Item, dataService DataService) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			for id := range idChan {
				func(id2 int) {
					item, err := dataService.getItem(id2)
					if err != nil {
						log.Print(errors.Wrap(err, "Fan out ids: "))
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
