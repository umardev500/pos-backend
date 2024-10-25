package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
	"github.com/umardev500/pos-backend/models"
)

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) contracts.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create creates a new user
func (u *userRepository) Create(ctx context.Context, payload models.CreateUserRequest) error {
	query := u.db.GetConn(ctx)

	sql := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err := query.Exec(ctx, sql, payload.Username, payload.Email, payload.Password)
	return err
}

// Delete deletes a user or multiple users
func (u *userRepository) Delete(ctx context.Context, ids []uuid.UUID) error {
	query := u.db.GetConn(ctx)

	sql := `DELETE FROM users WHERE id = ANY($1)`
	_, err := query.Exec(ctx, sql, ids)
	return err
}

// Find finds all users
func (u *userRepository) Find(ctx context.Context) ([]models.User, error) {
	query := u.db.GetConn(ctx)

	sql := `SELECT * FROM users`
	rows, err := query.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
