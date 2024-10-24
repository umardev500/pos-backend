package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/models"
)

type userHandler struct {
	usecase contracts.UserUsecaseInterface
}

func NewUserHandler(usecase contracts.UserUsecaseInterface) contracts.UserHandlerInterface {
	return &userHandler{
		usecase: usecase,
	}
}

func (u *userHandler) Create(c *fiber.Ctx) error {
	var payload models.CreateUserRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	return c.SendStatus(200)
}
