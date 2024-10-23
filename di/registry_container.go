package di

import "github.com/umardev500/pos-backend/contracts"

func NewRegistryContainer() []contracts.ContainerInterface {
	return []contracts.ContainerInterface{
		NewUserContainer(),
	}
}
