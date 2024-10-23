package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/app/handler"
	"github.com/umardev500/pos-backend/app/repository"
	"github.com/umardev500/pos-backend/app/usecase"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
)

type UserContainer struct {
	handler contracts.UserHandlerInterface
}

func NewUserContainer(db *database.DB, v contracts.ValidatorInterface) contracts.ContainerInterface {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase, v)

	return &UserContainer{
		handler: handler,
	}
}

func (u *UserContainer) RegisterApi(router fiber.Router) {
	r := router.Group("user")
	r.Post("/", u.handler.Create)
}

func (u *UserContainer) RegisterWeb(router fiber.Router) {}
