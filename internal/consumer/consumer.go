package consumer

import (
	"context"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/publisher"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Consumer is an interface to for the consumption of the queue, returning a id channel
type Consumer interface {
	Consume(ctx context.Context) (<-chan amqp.Delivery, error)
}

// Run handles the execution of the consumer, firing off workers concurrently
func Run(ctx context.Context, cfg *models.Config, consumer Consumer, provider publisher.Provider, writer Writer) error {
	idChan, err := consumer.Consume(ctx)
	if err != nil {
		return errors.Wrap(err, "consuming idChan ")
	}

	w := Worker{provider: provider, writer: writer, idChan: idChan}

	var wg sync.WaitGroup

	for i := 0; i < cfg.WorkerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			w.Run(ctx)
		}()

	}

	wg.Wait()

	return nil
}
