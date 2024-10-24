package models

import "github.com/umardev500/pos-backend/constants"

type ErrItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Case    string `json:"case"`
}

type Response struct {
	Code      int                   `json:"-"`
	Success   bool                  `json:"success"`
	Message   constants.MessageText `json:"message,omitempty"`
	Data      interface{}           `json:"data,omitempty"`
	ErrorCode constants.ErrorCode   `json:"error_code,omitempty"`
	Errors    interface{}           `json:"errors,omitempty"`
	RefCode   string                `json:"ref_code,omitempty"`
}
