package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
)

func NewRegistryContainer(db *database.DB, v *validator.Validate) []contracts.ContainerInterface {
	return []contracts.ContainerInterface{
		NewUserContainer(db, v),
	}
}
