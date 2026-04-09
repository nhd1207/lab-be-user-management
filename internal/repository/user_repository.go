package repository

import "github.com/nhd1207/be-user-management/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
}
