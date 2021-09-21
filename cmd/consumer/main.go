package main

import "github.com/JakeHumphries/gymshark-hacker-news/internal/consumer"

func main() {
	consumer.Consume(consumer.HttpService{})
}