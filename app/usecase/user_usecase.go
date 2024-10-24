package usecase

import "github.com/umardev500/pos-backend/contracts"

type userUsecase struct {
	repo     contracts.UserRepositoryInterface
	validate contracts.ValidatorInterface
}

func NewUserUsecase(repo contracts.UserRepositoryInterface, v contracts.ValidatorInterface) contracts.UserUsecaseInterface {
	return &userUsecase{
		repo:     repo,
		validate: v,
	}
}
