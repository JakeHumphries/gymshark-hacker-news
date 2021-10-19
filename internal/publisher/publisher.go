package publisher

import (
	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Publisher interface {
	Publish(id int) error
}

// Provider is an interface for getting hackernews data
type Provider interface {
	GetTopStories() ([]int, error)
	GetItem(id int) (*models.Item, error)
}

func Run(pub Publisher, prov Provider) error {
	ids, err := prov.GetTopStories()
	if err != nil {
		log.Fatal(errors.Wrap(err, "execute"))
	}

	for _, id := range ids {
		if err := pub.Publish(id); err != nil {
			log.Print(errors.Wrap(err, "publisher"))
		}
	}

	return nil
}
