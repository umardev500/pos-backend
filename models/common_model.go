package models

import "github.com/google/uuid"

type ListOfID struct {
	IDs []uuid.UUID `json:"ids"`
}
