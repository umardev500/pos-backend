package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
)

type UserContainer struct{}

func NewUserContainer() contracts.ContainerInterface {
	return &UserContainer{}
}

func (u *UserContainer) RegisterApi(router fiber.Router) {}
func (u *UserContainer) RegisterWeb(router fiber.Router) {}
