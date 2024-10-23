package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
)

type userHandler struct {
	usecase contracts.UserUsecaseInterface
}

func NewUserHandler(usecase contracts.UserUsecaseInterface) contracts.UserHandlerInterface {
	return &userHandler{
		usecase: usecase,
	}
}

func (u *userHandler) Create(ctx *fiber.Ctx) error {
	return nil
}
