package repository

import (
	"context"

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

func (u *userRepository) Create(ctx context.Context, payload models.CreateUserRequest) error {
	query := u.db.GetConn(ctx)

	sql := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err := query.Exec(ctx, sql, payload.Username, payload.Email, payload.Password)
	return err
}
