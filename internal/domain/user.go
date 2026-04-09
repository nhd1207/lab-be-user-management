package domain

import (
	"errors"
	"time"
)

type User struct {
	ID           string
	Email        string
	Username     string
	PasswordHash string
	Provider     AuthProvider
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AuthProvider string

const (
	ProviderLocal     AuthProvider = "local"
	ProviderGoogle    AuthProvider = "google"
	ProviderMicrosoft AuthProvider = "microsoft"
)

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Username == "" {
		return errors.New("Username is required")
	}
	return nil
}
