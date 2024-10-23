package repository

import (
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
)

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) contracts.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
