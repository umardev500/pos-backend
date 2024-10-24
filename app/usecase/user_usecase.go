package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos-backend/constants"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/models"
	"github.com/umardev500/pos-backend/utils"
)

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

// Create creates a new user
func (u *userUsecase) Create(ctx context.Context, payload models.CreateUserRequest) models.Response {
	validationErrs := u.validate.Struct(payload)
	if len(validationErrs) > 0 {
		return models.Response{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: constants.MessageCommonValidationFailed,
			Data:    validationErrs,
		}
	}

	// Hash password
	hash, err := utils.HashPassword(payload.Password)
	if err != nil {
		refCode := utils.LogError(err)

		return models.Response{
			Code:    fiber.StatusInternalServerError,
			Success: false,
			Message: constants.MessageCommonInternalServerError,
			RefCode: refCode,
		}
	}

	payload.Password = hash

	err = u.repo.Create(ctx, payload)
	if err != nil {
		refCode := utils.LogError(err)

		return models.Response{
			Code:    fiber.StatusInternalServerError,
			Success: false,
			Message: constants.MessageUserCreateFailed,
			RefCode: refCode,
		}
	}

	return models.Response{
		Code:    fiber.StatusCreated,
		Success: true,
		Message: constants.MessageUserCreateSuccess,
	}
}
