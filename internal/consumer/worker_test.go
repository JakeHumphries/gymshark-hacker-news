package consumer

import (
	"context"
	"testing"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
	m.Called(ctx, item)
	return nil, nil
}

func (m *MockWriter) GetAllItems(ctx context.Context) ([]models.Item, error) {
	m.Called(ctx)
	return nil, nil
}

type MockProvider struct {
	mock.Mock
}

func (m *MockProvider) GetTopStories() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m *MockProvider) GetItem(id int) (*models.Item, error) {
	args := m.Called(id)

	itemArg, ok := args.Get(0).(*models.Item)
	if !ok {
		return nil, nil
	}

	return itemArg, args.Error(1)
}

func TestConsumer_Run(t *testing.T) {
	tests := []struct {
		name       string
		writer     *MockWriter
		provider   *MockProvider
		expected   func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider)
		assertions func(writerMock *MockWriter)
	}{
		{
			name:     "calls saveItem for every item returned from provider",
			writer:   &MockWriter{},
			provider: &MockProvider{},
			expected: func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider) {
				writerMock.On("SaveItem").Return(nil)
				providerMock.On("GetItem", 1).Return(&models.Item{Id: 1}, nil)
				providerMock.On("GetItem", 2).Return(&models.Item{Id: 2}, nil)
				providerMock.On("GetItem", 3).Return(&models.Item{Id: 3}, nil)

			},
			assertions: func(writerMock *MockWriter) {
				writerMock.AssertNumberOfCalls(t, "SaveItem", 3)
			},
		},
		{
			name:     "doesnt call saveItem when item is set to deleted or dead",
			writer:   &MockWriter{},
			provider: &MockProvider{},
			expected: func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider) {
				writerMock.On("SaveItem").Return(nil)
				providerMock.On("GetItem", 1).Return(&models.Item{Id: 1, Dead: true}, nil)
				providerMock.On("GetItem", 2).Return(&models.Item{Id: 2, Deleted: true}, nil)
				providerMock.On("GetItem", 3).Return(&models.Item{Id: 3, Dead: true, Deleted: true}, nil)
			},
			assertions: func(writerMock *MockWriter) {
				writerMock.AssertNumberOfCalls(t, "SaveItem", 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != nil {
				tt.expected(t, tt.writer, tt.provider)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			
			idChan := make(chan int)

			var id = 1
			for i := 0; i < 3; i++ {
				idChan <- id
				id++
			}

			w := Worker{tt.provider, tt.writer}

			w.Run(ctx, idChan)

			if tt.assertions != nil {
				tt.assertions(tt.writer)
			}
		})
	}
}