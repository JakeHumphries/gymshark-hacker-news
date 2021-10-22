package queue

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Queue holds the logic for interacting with a rabbitMQ queue
type Queue struct {
	ch   *amqp.Channel
	conn *amqp.Connection
}

// New creates a new rabbitMQ Queue
func New(ctx context.Context, cfg models.Config) (*Queue, error) {
	conn, err := connect(ctx, cfg)
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
		"", // exchange
		"ItemQueue", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%d", id)),
		})
}

func (q Queue) Consume(idChan chan int) error {
	msgs, err := q.ch.Consume("ItemQueue", "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for msg := range msgs {
		id, err := strconv.Atoi(string(msg.Body))
		if err != nil {
			log.Print(errors.Wrap(err, "worker"))
		}

		idChan <- id
	}

	return nil
}
