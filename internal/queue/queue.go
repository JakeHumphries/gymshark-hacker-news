package queue

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// Queue holds the logic for interacting with a rabbitMQ queue
type Queue struct {
	ch   *amqp.Channel
	conn *amqp.Connection
}

// New creates a new rabbitMQ Queue
func New(ctx context.Context) (*Queue, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to RabbitMQ ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open a channel ")
	}

	_, err = ch.QueueDeclare("ItemQueue", false, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to declare a queue ")
	}

	return &Queue{ch, conn}, nil
}

func (q Queue) Close() error {
	if err := q.conn.Close(); err != nil {
		return errors.Wrap(err, "failed to close the connection")
	}

	if err := q.ch.Close(); err != nil {
		return errors.Wrap(err, "failed to close the channel")
	}

	return nil
}

func (q Queue) Publish(id int) error {
	return q.ch.Publish(
		"",
		"ItemQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%d", id)),
		})
}

func (q Queue) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	msgs, err := q.ch.Consume("ItemQueue", "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
