package handler

import (
	"github.com/go-playground/validator/v10"
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

func (u *userHandler) Create(ctx *fiber.Ctx) error {
	var payload models.CreateUserRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	valiato := validator.New()
	if err := valiato.Struct(payload); err != nil {
		return err
	}

	return nil
}
