package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nhd1207/be-user-management/internal/domain"
)

// mock service
type MockUserService struct{}

func (m *MockUserService) CreateUser(email, username, password string) (*domain.User, error) {
	return &domain.User{
		ID:       "1",
		Email:    email,
		Username: username,
	}, nil
}

func TestCreateUserHandler(t *testing.T) {
	s := &MockUserService{}
	h := NewUserHandler(s)

	body := map[string]string{
		"email":    "test@example.com",
		"username": "testuser",
		"password": "password",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.CreateUser(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}
