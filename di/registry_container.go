package di

import (
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
)

func NewRegistryContainer(db *database.DB, v contracts.ValidatorInterface) []contracts.ContainerInterface {
	return []contracts.ContainerInterface{
		NewUserContainer(db, v),
	}
}
