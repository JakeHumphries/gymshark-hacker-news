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
	m.Called()
	return nil, nil
}

func (m *MockWriter) GetAllItems(ctx context.Context) ([]models.Item, error) {
	m.Called()
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

type MockConsumer struct {
	mock.Mock
}

func (m *MockConsumer) Consume(idChan chan int) error {
	m.Called()

	var id = 1
	for i := 0; i < 3; i++ {
		idChan <- id
		id++
	}
	return nil
}

func TestConsumer_Run(t *testing.T) {

	tests := []struct {
		name       string
		writer     *MockWriter
		provider   *MockProvider
		consumer   *MockConsumer
		expected   func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider, consumerMock *MockConsumer)
		assertions func(writerMock *MockWriter)
	}{
		{
			name:     "calls saveItem for every item returned from provider",
			writer:   &MockWriter{},
			provider: &MockProvider{},
			consumer: &MockConsumer{},
			expected: func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider, consumerMock *MockConsumer) {
				writerMock.On("SaveItem").Return(nil)
				consumerMock.On("Consume").Return(nil)
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
			consumer: &MockConsumer{},
			expected: func(t *testing.T, writerMock *MockWriter, providerMock *MockProvider, consumerMock *MockConsumer) {
				writerMock.On("SaveItem").Return(nil)
				consumerMock.On("Consume").Return(nil)
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
				tt.expected(t, tt.writer, tt.provider, tt.consumer)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			cfg := models.Config{
				DatabaseName:     "testName",
				DatabaseUser:     "testUser",
				DatabasePassword: "testPass",
				DatabasePort:     "testPort",
				Cron:             "0 30 * * * *",
				WorkerCount:      10,
			}

			Run(ctx, &cfg, tt.consumer, tt.provider, tt.writer)

			if tt.assertions != nil {
				tt.assertions(tt.writer)
			}
		})
	}
}
