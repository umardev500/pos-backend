package contracts

import "github.com/umardev500/pos-backend/models"

type ValidatorInterface interface {
	Struct(payload interface{}) []models.ErrItem
}
