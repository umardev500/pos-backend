package contracts

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos-backend/models"
)

type UserHandlerInterface interface {
	Create(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

type UserUsecaseInterface interface {
	Create(context.Context, models.CreateUserRequest) models.Response
	Delete(context.Context, []uuid.UUID) models.Response
}

type UserRepositoryInterface interface {
	Create(context.Context, models.CreateUserRequest) error
	Delete(context.Context, []uuid.UUID) error
	Find(context.Context) ([]models.User, error)
}
