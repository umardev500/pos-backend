package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/contracts"
)

type Router struct {
	container []contracts.ContainerInterface
}

func NewRouter(container []contracts.ContainerInterface) contracts.RouterInterface {
	return &Router{
		container: container,
	}
}

func (r *Router) Setup(app *fiber.App) {
	// Register API routes
	api := app.Group("/api")
	r.setupApi(api)

	// Register Web routes
	web := app.Group("/")
	r.setupWeb(web)
}

// setupApi defines the api routes
func (r *Router) setupApi(router fiber.Router) {
	for _, container := range r.container {
		container.RegisterApi(router)
	}
}

// setupWeb defines the web routes
func (r *Router) setupWeb(router fiber.Router) {
	for _, container := range r.container {
		container.RegisterWeb(router)
	}
}
