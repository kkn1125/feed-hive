package handler

import (
	"bytes"
	"encoding/json"
	"feedhive/users/model"
	"feedhive/users/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() (*[]model.User, error) {
	args := m.Called()
	return args.Get(0).(*[]model.User), args.Error(1)
}

func (m *MockUserRepository) FindById(id string) (*model.User, error) {
	args := m.Called()
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called()
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *model.User) (uint, error) {
	args := m.Called(user)
	return args.Get(0).(uint), args.Error(1)
}

func TestControllerSuite(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockUserRepository)
	userHandler := NewUserHandler(mockRepo)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	hashedPassword := util.HashPassword("1234")

	requestBody := map[string]interface{}{
		"name":     "devkimson",
		"email":    "chaplet01@gmail.com",
		"password": hashedPassword,
	}

	jsonValue, _ := json.Marshal(requestBody)
	c.Request, _ = http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(uint(1), nil)

	userHandler.CreateUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}

	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(0), response["created"])
	// assert.Equal(t, "chaplet01@gmail.com", response["email"])

	mockRepo.AssertExpectations(t)
}
