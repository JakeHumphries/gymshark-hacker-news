package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := consumer.Connect(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(mongoClient)

	consumer.Consume(consumer.HttpService{})
}