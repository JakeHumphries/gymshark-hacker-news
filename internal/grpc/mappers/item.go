package grpc

import (
	"github.com/JakeHumphries/gymshark-hacker-news/internal/grpc/protobufs"
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
)

func ToProto(item models.Item) *protobufs.Item {
	return &protobufs.Item{
		By: item.By,
		Title: item.Title,
		Url: item.Url,
		Text: item.Text,
		Descendants: item.Descendants,
		Id: item.Id,
		Score: item.Score,
		Time: item.Time,
		Poll: item.Poll,
		Kids: item.Kids,
		Parts: item.Parts,
		Deleted: item.Deleted,
		Dead: item.Dead,
	}
}