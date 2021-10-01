package consumer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
	m.Called()
	return nil, nil
}

type MockItemProvider struct {
	mock.Mock
}

func (m *MockItemProvider) GetTopStories() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}

func (m *MockItemProvider) GetItem(id int) (*models.Item, error) {
	item := models.Item{
		Deleted: false,
		Dead:    false,
	}

	return &item, nil
}

func TestConsumer_Execute(t *testing.T) {
	fmt.Println("hello world")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println(ctx)

	cfg := models.Config{
		DatabaseName:     "testName",
		DatabaseUser:     "testUser",
		DatabasePassword: "testPass",
		DatabasePort:     "testPort",
		Cron:             "0 30 * * * *",
		WorkerCount:      10,
	}

	fmt.Println(cfg)

	itemRepoMock := new(MockItemRepository)
	itemProviderMock := new(MockItemProvider)

	itemRepoMock.On("SaveItem").Return(nil)

	Execute(ctx, cfg, itemRepoMock, itemProviderMock)

	itemRepoMock.AssertNumberOfCalls(t, "SaveItem", 5)
	itemRepoMock.AssertExpectations(t)
}
