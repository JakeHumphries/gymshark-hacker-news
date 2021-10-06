package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReader struct {
	mock.Mock
}

func (m *MockReader) GetAllItems(ctx context.Context) ([]models.Item, error) {
	args := m.Called()

	itemArg, ok := args.Get(0).(*models.Item)
	if !ok {
		return nil, args.Error(0)
	}

	items := []models.Item{*itemArg}

	return items, args.Error(1)
}

func TestConsumer_GetAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tests := []struct {
		name       string
		mockReader *MockReader
		code       int
		result     []models.Item
		expected   func(t *testing.T, mockReader *MockReader)
		assertions func(mockReader *MockReader)
	}{
		{
			name:       "returns correct status and response on successful request",
			mockReader: &MockReader{},
			code:       http.StatusOK,
			result:     []models.Item{{Id: 1}},
			expected: func(t *testing.T, mockReader *MockReader) {
				mockReader.On("GetAllItems").Return(&models.Item{Id: 1}, nil)

			},
			assertions: func(mockReader *MockReader) {
				mockReader.AssertNumberOfCalls(t, "GetAllItems", 1)
			},
		},
		{
			name:       "returns correct status and response on failed request",
			mockReader: &MockReader{},
			code:       http.StatusInternalServerError,
			result:     nil,
			expected: func(t *testing.T, mockReader *MockReader) {
				mockReader.On("GetAllItems").Return(errors.New("test fail"))

			},
			assertions: func(mockReader *MockReader) {
				mockReader.AssertNumberOfCalls(t, "GetAllItems", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != nil {
				tt.expected(t, tt.mockReader)
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/all")

			a := New(tt.mockReader, ctx)
			a.GetAll(c)

			if tt.assertions != nil {
				tt.assertions(tt.mockReader)
			}

			assert.Equal(t, tt.code, rec.Code)

			if tt.code == http.StatusOK {
				responseData, err := ioutil.ReadAll(rec.Body)
				if err != nil {
					t.Fail()
					t.Logf("failed reading response body: %d", err)
				}

				var items []models.Item
				err = json.Unmarshal(responseData, &items)
				if err != nil {
					t.Fail()
					t.Logf("failed unmarshalling response body: %d", err)
				}

				assert.Equal(t, tt.result, items)
			}

		})
	}
}
