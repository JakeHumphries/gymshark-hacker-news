package api

import (
	"context"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

// CacheReader calls an instance of redis to create and interact with a cache
type CacheReader struct {
	itemReader ItemReader
	cache      *cache.Cache
	cfg        models.Config
}

// NewCacheReader creates a new cache reader
func NewCacheReader(itemReader ItemReader, cfg models.Config) *CacheReader {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost,
	})

	itemCache := cache.New(&cache.Options{
		Redis: client,
	})

	return &CacheReader{
		itemReader: itemReader,
		cache:      itemCache,
	}
}

func (c CacheReader) GetAllItems(ctx context.Context) ([]models.Item, error) {
	var items []models.Item

	err := c.cache.Once(&cache.Item{
		Key:   "all",
		Value: &items,
		TTL:   c.cfg.CacheTimout,
		Do: func(*cache.Item) (interface{}, error) {
			i, err := c.itemReader.GetAllItems(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "cache get all items")
			}
			return i, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}