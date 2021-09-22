package consumer

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type DbRepoMock struct {
	mock.Mock
}

func (m *DbRepoMock) SaveItem(item Item) error {
	m.Called()
	return nil
}

type DataServiceMock struct {
	mock.Mock
}

func (m *DataServiceMock) getTopStories() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}

func (m *DataServiceMock) getItem(id int) (*Item, error) {
	item := Item{
    Deleted: false,
    Dead: false,
	}

	return &item, nil
}

func TestConsumer_Consume(t *testing.T) {

	dbrepoMock := new(DbRepoMock)
	dataServiceMock := new(DataServiceMock)

	dbrepoMock.On("SaveItem").Return(nil) 

	Consume(dbrepoMock, dataServiceMock)

	dbrepoMock.AssertNumberOfCalls(t, "SaveItem", 5)
	dbrepoMock.AssertExpectations(t)

}
