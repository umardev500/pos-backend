package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/umardev500/pos-backend/constants"
	"github.com/umardev500/pos-backend/models"
)

// handleUniqueViolationError handles unique constraint error
func handleUniqueViolationError(pgErr *pgconn.PgError) *models.ErrItem {
	var err models.ErrItem = models.ErrItem{
		Case: "conflict",
	}

	field := extractFieldName(pgErr.Detail)
	jsonTag := getJsonTagForField(models.CreateUserRequest{}, field)
	err.Field = jsonTag
	err.Message = fmt.Sprintf("%s already exists", field)

	return &err
}

// extractFieldName extracts the field name from the error message
func extractFieldName(errMessage string) string {
	re := regexp.MustCompile(`Key \(([^)]+)\)=`)
	matches := re.FindStringSubmatch(errMessage)
	return matches[1]
}

// getJsonTagForField returns the json tag for the given field
func getJsonTagForField(structType interface{}, fieldName string) string {
	val := reflect.ValueOf(structType)
	typ := val.Type()

	// Ensure we're passing a struct type
	if typ.Kind() != reflect.Struct {
		return "unknown field"
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == fieldName {
			jsonTag := field.Tag.Get("json")
			return jsonTag
		}
	}

	return "unknown field"
}

// HandlePostgresError handles postgres error
func HandlePostgresError(err error) (interface{}, *constants.ErrorCode, constants.MessageText, int) {
	// var validationErrs models.ErrItem
	var statusCode = fiber.StatusInternalServerError
	var errCode constants.ErrorCode
	if err == nil {
		return nil, nil, "", statusCode
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// Unique constraint error
		case "23505":
			errItem := handleUniqueViolationError(pgErr)
			errCode = constants.ConflictErrorType
			return errItem, &errCode, "", fiber.StatusConflict
		default:
			errCode = constants.InternalServerErrorType
			return nil, &errCode, constants.MessageCommonInternalServerError, fiber.StatusInternalServerError
		}
	}

	return nil, nil, "", fiber.StatusOK
}
