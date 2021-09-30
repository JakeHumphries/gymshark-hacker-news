package consumer

import (
	"context"
	"testing"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockItemSaver struct {
	mock.Mock
}

func (m *MockItemSaver) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
	m.Called()
	return nil, nil
}

type MockDataGetter struct {
	mock.Mock
}

func (m *MockDataGetter) GetTopStories() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}

func (m *MockDataGetter) GetItem(id int) (*models.Item, error) {
	item := models.Item{
		Deleted: false,
		Dead:    false,
	}

	return &item, nil
}

func TestConsumer_Execute(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbrepoMock := new(MockItemSaver)
	dataServiceMock := new(MockDataGetter)

	dbrepoMock.On("SaveItem").Return(nil)

	Execute(ctx, dbrepoMock, dataServiceMock)

	dbrepoMock.AssertNumberOfCalls(t, "SaveItem", 5)
	dbrepoMock.AssertExpectations(t)
}
