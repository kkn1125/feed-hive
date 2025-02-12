package handler

import (
	"bytes"
	"encoding/json"
	"feedhive/feeds/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFeedRepository struct {
	mock.Mock
}

func (m *MockFeedRepository) FindAll() (*[]model.Feed, error) {
	args := m.Called()
	return args.Get(0).(*[]model.Feed), args.Error(1)
}

func (m *MockFeedRepository) FindById(id string) (*model.Feed, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Feed), args.Error(1)
}

func (m *MockFeedRepository) FindNotificationById(id string) (*model.Notification, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Notification), args.Error(1)
}

func (m *MockFeedRepository) Create(feed *model.Feed) (uint, error) {
	args := m.Called(feed)
	feed.ID++
	log.Printf("✨ 확인 feed: %v", feed)
	log.Printf("✨ 확인 args: %v", args)
	return args.Get(0).(uint), args.Error(1)
}

func TestSuite(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockFeedRepository)
	feedHandler := NewFeedHandler(mockRepo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := map[string]interface{}{
		"user_id": 1,
		"content": "test",
	}

	jsonValue, _ := json.Marshal(requestBody)
	c.Request, _ = http.NewRequest("POST", "/api/feeds", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	mockRepo.On("Create", mock.AnythingOfType("*model.Feed")).Return(uint(1), nil)
	feedHandler.CreateFeed(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	log.Printf("✨ 확인 response: %v", response)
	assert.Equal(t, uint(1), uint(response["created"].(float64)))
}
