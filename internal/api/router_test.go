package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JakeHumphries/gymshark-hacker-news/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func setupRequest(path string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(path)

	return rec, c
}

func decodeRequest(t *testing.T, body *bytes.Buffer) []models.Item {
	responseData, err := ioutil.ReadAll(body)
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

	return items
}

func (m *MockHandler) GetAllItems(ctx context.Context) ([]models.Item, error) {
	args := m.Called()

	itemArg, ok := args.Get(0).([]models.Item)
	if !ok {
		return nil, args.Error(0)
	}

	return itemArg, args.Error(1)
}

func (m *MockHandler) GetStories(ctx context.Context) ([]models.Item, error) {
	args := m.Called()

	itemArg, ok := args.Get(0).([]models.Item)
	if !ok {
		return nil, args.Error(0)
	}

	return itemArg, args.Error(1)
}

func (m *MockHandler) GetJobs(ctx context.Context) ([]models.Item, error) {
	args := m.Called()

	itemArg, ok := args.Get(0).([]models.Item)
	if !ok {
		return nil, args.Error(0)
	}

	return itemArg, args.Error(1)
}

func TestConsumer_GetAll(t *testing.T) {
	tests := []struct {
		name        string
		mockHandler *MockHandler
		code        int
		result      []models.Item
		expected    func(t *testing.T, mockHandler *MockHandler)
		assertions  func(mockHandler *MockHandler)
	}{
		{
			name:        "returns correct status and response on successful request",
			mockHandler: &MockHandler{},
			code:        http.StatusOK,
			result:      []models.Item{{Id: 1}},
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetAllItems").Return([]models.Item{{Id: 1}}, nil)

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetAllItems", 1)
			},
		},
		{
			name:        "returns correct status and response on failed request",
			mockHandler: &MockHandler{},
			code:        http.StatusInternalServerError,
			result:      nil,
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetAllItems").Return(errors.New("test fail"))

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetAllItems", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != nil {
				tt.expected(t, tt.mockHandler)
			}

			rec, c := setupRequest("/all")

			r := NewRouter(tt.mockHandler)

			r.All(c)

			if tt.assertions != nil {
				tt.assertions(tt.mockHandler)
			}

			assert.Equal(t, tt.code, rec.Code)

			if tt.code == http.StatusOK {
				items := decodeRequest(t, rec.Body)

				assert.Equal(t, tt.result, items)
			}
		})
	}
}

func TestConsumer_GetStories(t *testing.T) {
	tests := []struct {
		name        string
		mockHandler *MockHandler
		code        int
		result      []models.Item
		expected    func(t *testing.T, mockHandler *MockHandler)
		assertions  func(mockHandler *MockHandler)
	}{
		{
			name:        "returns correct status and response on successful request",
			mockHandler: &MockHandler{},
			code:        http.StatusOK,
			result:      []models.Item{{Id: 1, Type: "story"}},
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetStories").Return([]models.Item{{Id: 1, Type: "story"}}, nil)

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetStories", 1)
			},
		},
		{
			name:        "returns correct status and response on failed request",
			mockHandler: &MockHandler{},
			code:        http.StatusInternalServerError,
			result:      nil,
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetStories").Return(errors.New("test fail"))

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetStories", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != nil {
				tt.expected(t, tt.mockHandler)
			}

			rec, c := setupRequest("/stories")

			r := NewRouter(tt.mockHandler)

			r.Stories(c)

			if tt.assertions != nil {
				tt.assertions(tt.mockHandler)
			}

			assert.Equal(t, tt.code, rec.Code)

			if tt.code == http.StatusOK {
				items := decodeRequest(t, rec.Body)

				assert.Equal(t, tt.result, items)
			}
		})
	}
}

func TestConsumer_GetJobs(t *testing.T) {
	tests := []struct {
		name        string
		mockHandler *MockHandler
		code        int
		result      []models.Item
		expected    func(t *testing.T, mockHandler *MockHandler)
		assertions  func(mockHandler *MockHandler)
	}{
		{
			name:        "returns correct status and response on successful request",
			mockHandler: &MockHandler{},
			code:        http.StatusOK,
			result:      []models.Item{{Id: 1, Type: "job"}},
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetJobs").Return([]models.Item{{Id: 1, Type: "job"}}, nil)

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetJobs", 1)
			},
		},
		{
			name:        "returns correct status and response on failed request",
			mockHandler: &MockHandler{},
			code:        http.StatusInternalServerError,
			result:      nil,
			expected: func(t *testing.T, mockHandler *MockHandler) {
				mockHandler.On("GetJobs").Return(errors.New("test fail"))

			},
			assertions: func(mockHandler *MockHandler) {
				mockHandler.AssertNumberOfCalls(t, "GetJobs", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != nil {
				tt.expected(t, tt.mockHandler)
			}

			rec, c := setupRequest("/jobs")

			r := NewRouter(tt.mockHandler)

			r.Jobs(c)

			if tt.assertions != nil {
				tt.assertions(tt.mockHandler)
			}

			assert.Equal(t, tt.code, rec.Code)

			if tt.code == http.StatusOK {
				items := decodeRequest(t, rec.Body)

				assert.Equal(t, tt.result, items)
			}
		})
	}
}
