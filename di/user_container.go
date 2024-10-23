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

func NewUserContainer(db *database.DB) contracts.ContainerInterface {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase)

	return &UserContainer{
		handler: handler,
	}
}

func (u *UserContainer) RegisterApi(router fiber.Router) {}

func (u *UserContainer) RegisterWeb(router fiber.Router) {}
