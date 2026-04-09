package service

import (
	"testing"

	"github.com/nhd1207/be-user-management/internal/domain"
)

type MockUserRepo struct {
	users map[string]*domain.User
}

func (m *MockUserRepo) Create(user *domain.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *MockUserRepo) GetByEmail(email string) (*domain.User, error) {
	user, ok := m.users[email]

	if !ok {
		return nil, nil
	}
	return user, nil
}

type MockHasher struct{}

func (m *MockHasher) Hash(password string) (string, error) {
	return "hashed_" + password, nil
}

func (m *MockHasher) Compare(hash, password string) error {
	return nil
}

func TestCreateUser_Success(t *testing.T) {
	repo := &MockUserRepo{users: map[string]*domain.User{}}
	hasher := &MockHasher{}
	service := NewUserService(repo, hasher)

	user, err := service.CreateUser("test@example.com", "testuser", "password")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if user.PasswordHash != "hashed_password" {
		t.Errorf("Expected password hash to be 'hashed_password', got %s", user.PasswordHash)
	}
}

func TestCreateUser_Duplicate(t *testing.T) {
	repo := &MockUserRepo{
		users: map[string]*domain.User{
			"test@example.com": {},
		},
	}

	hasher := &MockHasher{}
	service := NewUserService(repo, hasher)

	_, err := service.CreateUser("test@example.com", "testuser", "password")

	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}

	if err.Error() != "User already exists" {
		t.Errorf("Expected error message 'User already exists', got %s", err.Error())
	}
}
