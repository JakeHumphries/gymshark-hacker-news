package queue

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func connect(ctx context.Context) (*amqp.Connection, error) {
	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to RabbitMQ: %v %w", ctx.Err(), err)
		default:
			conn, err := amqp.Dial("amqp://guest:guest@gymshark-hacker-news-rabbitmq-1/")
			if err == nil {
				log.Print("rabbitmq is now connected")
				return conn, nil
			}
		}
	}
}
