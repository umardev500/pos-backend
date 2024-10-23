package usecase

import "github.com/umardev500/pos-backend/contracts"

type userUsecase struct {
	repo contracts.UserRepositoryInterface
}

func NewUserUsecase(repo contracts.UserRepositoryInterface) contracts.UserUsecaseInterface {
	return &userUsecase{
		repo: repo,
	}
}
