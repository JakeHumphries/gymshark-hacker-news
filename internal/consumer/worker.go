package consumer

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func worker(ctx context.Context, idChan chan int, itemProvider ItemProvider, itemRepository ItemRepository, wg sync.WaitGroup) {
	for id := range idChan {
		item, err := itemProvider.GetItem(id)
		if err != nil {
			log.Print(errors.Wrap(err, "worker: "))
		} else if !item.Dead && !item.Deleted {
			_, err := itemRepository.SaveItem(ctx, *item)
			if err != nil {
				log.Print(errors.Wrap(err, "worker: "))
			}
		}
	}
	wg.Done()

}
