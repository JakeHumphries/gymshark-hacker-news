package consumer

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
// 	"github.com/stretchr/testify/mock"
// )

// type MockItemRepository struct {
// 	mock.Mock
// }

// func (m *MockItemRepository) SaveItem(ctx context.Context, item models.Item) (*models.Item, error) {
// 	m.Called()
// 	return nil, nil
// }

// func (m *MockItemRepository) GetAllItems(ctx context.Context) ([]models.Item, error) {
// 	m.Called()
// 	return nil, nil
// }

// type MockItemProvider struct {
// 	mock.Mock
// }

// func (m *MockItemProvider) GetTopStories() ([]int, error) {
// 	return []int{1, 2, 3}, nil
// }

// func (m *MockItemProvider) GetItem(id int) (*models.Item, error) {
// 	args := m.Called(id)

// 	itemArg, ok := args.Get(0).(*models.Item)
// 	if !ok {
// 		return nil, nil
// 	}

// 	return itemArg, args.Error(1)
// }
// func TestConsumer_Execute(t *testing.T) {

// 	tests := []struct {
// 		name           string
// 		itemRepository *MockItemRepository
// 		itemProvider   *MockItemProvider
// 		expected       func(t *testing.T, itemRepoMock *MockItemRepository, itemProviderMock *MockItemProvider)
// 		assertions     func(itemRepoMock *MockItemRepository)
// 	}{
// 		{
// 			name:           "calls saveItem for every item returned from provider",
// 			itemRepository: &MockItemRepository{},
// 			itemProvider:   &MockItemProvider{},
// 			expected: func(t *testing.T, itemRepoMock *MockItemRepository, itemProviderMock *MockItemProvider) {
// 				itemRepoMock.On("SaveItem").Return(nil)
// 				itemProviderMock.On("GetItem", 1).Return(&models.Item{Id: 1}, nil)
// 				itemProviderMock.On("GetItem", 2).Return(&models.Item{Id: 2}, nil)
// 				itemProviderMock.On("GetItem", 3).Return(&models.Item{Id: 3}, nil)

// 			},
// 			assertions: func(itemRepoMock *MockItemRepository) {
// 				itemRepoMock.AssertNumberOfCalls(t, "SaveItem", 3)
// 			},
// 		},
// 		{
// 			name:           "doesnt call saveItem when item is set to deleted or dead",
// 			itemRepository: &MockItemRepository{},
// 			itemProvider:   &MockItemProvider{},
// 			expected: func(t *testing.T, itemRepoMock *MockItemRepository, itemProviderMock *MockItemProvider) {
// 				itemRepoMock.On("SaveItem").Return(nil)
// 				itemProviderMock.On("GetItem", 1).Return(&models.Item{Id: 1, Dead: true}, nil)
// 				itemProviderMock.On("GetItem", 2).Return(&models.Item{Id: 2, Deleted: true}, nil)
// 				itemProviderMock.On("GetItem", 3).Return(&models.Item{Id: 3, Dead: true, Deleted: true}, nil)
// 			},
// 			assertions: func(itemRepoMock *MockItemRepository) {
// 				itemRepoMock.AssertNumberOfCalls(t, "SaveItem", 0)
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.expected != nil {
// 				tt.expected(t, tt.itemRepository, tt.itemProvider)
// 			}

// 			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 			defer cancel()

// 			cfg := models.Config{
// 				DatabaseName:     "testName",
// 				DatabaseUser:     "testUser",
// 				DatabasePassword: "testPass",
// 				DatabasePort:     "testPort",
// 				Cron:             "0 30 * * * *",
// 				WorkerCount:      10,
// 			}

// 			Execute(ctx, cfg, tt.itemRepository, tt.itemProvider)

// 			if tt.assertions != nil {
// 				tt.assertions(tt.itemRepository)
// 			}
// 		})
// 	}
// }

type TST struct{}
