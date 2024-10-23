package handler

import "github.com/umardev500/pos-backend/contracts"

type userHandler struct{}

func NewUserHandler() contracts.UserHandlerInterface {
	return &userHandler{}
}
