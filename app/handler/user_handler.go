package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/models"
	"github.com/umardev500/pos-backend/utils"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	respose := u.usecase.Create(ctx, payload)

	return c.Status(respose.Code).JSON(respose)
}

func (u *userHandler) Delete(c *fiber.Ctx) error {
	var payload models.ListOfID
	if err := c.BodyParser(&payload); err != nil {
		return utils.HandleBodyParserError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp := u.usecase.Delete(ctx, payload.IDs)

	return c.Status(resp.Code).JSON(resp)
}
