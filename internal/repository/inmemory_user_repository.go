package repository

import "github.com/nhd1207/be-user-management/internal/domain"

type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.users[user.Email] = user
	return nil
}

func (r *InMemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, nil
	}
	return u, nil
}
