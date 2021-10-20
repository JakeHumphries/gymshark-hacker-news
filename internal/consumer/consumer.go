package consumer

import (
	"context"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/publisher"
	"github.com/pkg/errors"
)

// Consumer is an interface to for the consumption of the queue, returning a id channel
type Consumer interface {
	Consume(idChan chan int) error
}

// Run handles the execution of the consumer, firing off workers concurrently
func Run(ctx context.Context, cfg *models.Config, consumer Consumer, provider publisher.Provider, writer Writer) error {

	var wg sync.WaitGroup

	idChan := make(chan int)

	w := Worker{provider: provider, writer: writer, idChan: idChan}

	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			w.Run(ctx)
		}()

	}

	if err := consumer.Consume(idChan); err != nil {
		return errors.Wrap(err, "consuming idChan ")
	}

	close(idChan)

	wg.Wait()

	return nil
}
