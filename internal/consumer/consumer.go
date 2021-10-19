package consumer

import (
	"context"
	"sync"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/publisher"
	"github.com/streadway/amqp"
)

type Consumer interface {
	Consume(ctx context.Context) (<-chan amqp.Delivery, error)
}

func Run(ctx context.Context, cfg *models.Config, consumer Consumer, provider publisher.Provider, writer Writer) {
	idChan, err := consumer.Consume(ctx)
	if err != nil {
		// do something with error
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
}
