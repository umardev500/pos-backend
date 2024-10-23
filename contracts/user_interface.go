package contracts

import "github.com/gofiber/fiber/v2"

type UserHandlerInterface interface {
	Create(*fiber.Ctx) error
}

type UserUsecaseInterface interface{}

type UserRepositoryInterface interface{}
