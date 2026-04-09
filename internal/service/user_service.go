package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nhd1207/be-user-management/internal/domain"
	"github.com/nhd1207/be-user-management/internal/repository"
)

type UserService struct {
	repo   repository.UserRepository
	hasher PasswordHasher
}

func NewUserService(repo repository.UserRepository, hasher PasswordHasher) *UserService {
	return &UserService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *UserService) CreateUser(email, username, password string) (*domain.User, error) {
	existing, _ := s.repo.GetByEmail(email)

	if existing != nil {
		return nil, errors.New("User already exists")
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           uuid.NewString(),
		Email:        email,
		Username:     username,
		PasswordHash: hash,
		Provider:     domain.ProviderLocal,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
