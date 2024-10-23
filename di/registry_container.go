package di

import (
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
)

func NewRegistryContainer(db *database.DB) []contracts.ContainerInterface {
	return []contracts.ContainerInterface{
		NewUserContainer(db),
	}
}
