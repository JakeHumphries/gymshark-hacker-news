package queue

import (
	"context"
	"fmt"
	"net/url"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func connect(ctx context.Context, cfg models.Config) (*amqp.Connection, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", ctx.Err())
		default:
			conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", url.QueryEscape(cfg.RabbitMqUser), url.QueryEscape(cfg.RabbitMqPassword), url.QueryEscape(cfg.RabbitMqHost)))
			if err == nil {
				log.Print("rabbitmq is now connected")
				return conn, nil
			}
		}
	}
}
