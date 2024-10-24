package contracts

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/models"
)

type UserHandlerInterface interface {
	Create(*fiber.Ctx) error
}

type UserUsecaseInterface interface {
	Create(context.Context, models.CreateUserRequest) models.Response
}

type UserRepositoryInterface interface {
	Create(context.Context, models.CreateUserRequest) error
}
