package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
)

type Router struct{}

func NewRouter() contracts.RouterInterface {
	return &Router{}
}

func (r *Router) Setup(router *fiber.App) {
	fmt.Println("hi")
}
