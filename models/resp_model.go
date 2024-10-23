package models

import "github.com/umardev500/pos-backend/constants"

type ValidationErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Case    string `json:"case"`
}

type Response struct {
	Code             int                 `json:"-"`
	Success          bool                `json:"success"`
	Message          string              `json:"message,omitempty"`
	Data             interface{}         `json:"data,omitempty"`
	ErrorCode        constants.ErrorCode `json:"error_code,omitempty"`
	ValidationErrors []ValidationErrors  `json:"validationErrors,omitempty"`
}
