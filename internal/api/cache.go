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
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": cfg.RedisHost,
		},
	})

	itemCache := cache.New(&cache.Options{
		Redis: ring,
	})

	return &CacheReader{
		itemReader: itemReader,
		cache:      itemCache,
	}
}

// GetAllItems will search the redis cache before calling the real item reader
func (c CacheReader) GetAllItems(ctx context.Context) ([]models.Item, error) {
	var items []models.Item

	err := c.cache.Once(&cache.Item{
		Key:   "all",
		Value: &items,
		TTL:   c.cfg.CacheTimout,
		Do: func(*cache.Item) (interface{}, error) {
			i, err := c.itemReader.GetAllItems(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "cache")
			}
			return i, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetStories will search the redis cache before calling the real item reader
func (c CacheReader) GetStories(ctx context.Context) ([]models.Item, error) {
	var items []models.Item

	err := c.cache.Once(&cache.Item{
		Key:   "stories",
		Value: &items,
		TTL:   c.cfg.CacheTimout,
		Do: func(*cache.Item) (interface{}, error) {
			i, err := c.itemReader.GetStories(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "cache")
			}
			return i, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetJobs will search the redis cache before calling the real item reader
func (c CacheReader) GetJobs(ctx context.Context) ([]models.Item, error) {
	var items []models.Item

	err := c.cache.Once(&cache.Item{
		Key:   "jobs",
		Value: &items,
		TTL:   c.cfg.CacheTimout,
		Do: func(*cache.Item) (interface{}, error) {
			i, err := c.itemReader.GetJobs(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "cache")
			}
			return i, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}
