package handler

import "github.com/umardev500/pos-backend/contracts"

type userHandler struct {
	usecase contracts.UserHandlerInterface
}

func NewUserHandler(usecase contracts.UserHandlerInterface) contracts.UserHandlerInterface {
	return &userHandler{
		usecase: usecase,
	}
}
