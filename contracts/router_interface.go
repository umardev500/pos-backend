package contracts

import "github.com/gofiber/fiber/v2"

type RouterInterface interface {
	Setup(*fiber.App)
}
