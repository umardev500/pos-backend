package contracts

import "github.com/gofiber/fiber/v2"

type ContainerInterface interface {
	RegisterApi(fiber.Router)
	RegisterWeb(fiber.Router)
}
