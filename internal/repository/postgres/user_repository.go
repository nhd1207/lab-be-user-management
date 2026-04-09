package postgres

import (
	"database/sql"

	"github.com/nhd1207/be-user-management/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `
	INSERT INTO users (id, email, username, password_hash, provider, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, user.ID, user.Email, user.Username, user.PasswordHash, user.Provider, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `
	SELECT id, email, username, password_hash, provider, created_at, updated_at
	FROM users
	WHERE email = $1
	`

	row := r.db.QueryRow(query, email)
	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.Provider,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
