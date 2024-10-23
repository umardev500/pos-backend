package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/constants"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/models"
)

type userHandler struct {
	usecase  contracts.UserUsecaseInterface
	validate contracts.ValidatorInterface
}

func NewUserHandler(usecase contracts.UserUsecaseInterface, v contracts.ValidatorInterface) contracts.UserHandlerInterface {
	return &userHandler{
		usecase:  usecase,
		validate: v,
	}
}

func (u *userHandler) Create(c *fiber.Ctx) error {
	var payload models.CreateUserRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validationErrs := u.validate.Struct(payload)

	return c.JSON(models.Response{
		Success:          false,
		Message:          "validation failed",
		ErrorCode:        constants.ValidationErrorType,
		ValidationErrors: validationErrs,
	})
}
