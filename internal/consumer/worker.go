package consumer

import (
	"log"
	"sync"

	"github.com/pkg/errors"
)

func worker(idChan chan int, itemProvider ItemProvider, itemRepository ItemRepository, wg sync.WaitGroup) {
	for id := range idChan {
			item, err := itemProvider.GetItem(id)
			if err != nil {
				log.Print(errors.Wrap(err, "worker: "))
			} else if !item.Dead && !item.Deleted {
				_, err := itemRepository.SaveItem(*item)
				if err != nil {
					log.Print(errors.Wrap(err, "worker: "))
				}
			}
	}
	wg.Done()

}