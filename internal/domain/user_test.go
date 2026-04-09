package domain

import "testing"

func TestUserValidate_Success(t *testing.T) {
	u := &User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	err := u.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUserValidate_MissingEmail(t *testing.T) {
	u := &User{
		Username: "testuser",
	}

	err := u.Validate()
	if err == nil {
		t.Error("Expected an error for missing email")
	}
}

func TestUserValidate_MissingUsername(t *testing.T) {
	u := &User{
		Email: "test@example.com",
	}

	err := u.Validate()
	if err == nil {
		t.Error("Expected an error for missing username")
	}
}
