package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos-backend/constants"
	"github.com/umardev500/pos-backend/models"
)

func HandleBodyParserError(c *fiber.Ctx, err error) error {

	switch {
	case uuid.IsInvalidLengthError(err):
		refCode := LogError(err)
		resp := models.Response{
			Code:      fiber.StatusUnprocessableEntity,
			Success:   false,
			Message:   constants.MessageCommonValidationFailed,
			ErrorCode: constants.ErrorCodeValidation,
			Errors: []models.ErrItem{
				{
					Field:   "uuid",
					Message: string(constants.MessageCommonUUIDInvalidLength),
					Case:    string(constants.CaseInvalidUUIDLength),
				},
			},
			RefCode: refCode,
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(resp)
	}

	return c.Status(fiber.StatusInternalServerError).JSON("")
}
