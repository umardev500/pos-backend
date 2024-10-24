package constants

type ErrorCode int

const (
	ValidationErrorType ErrorCode = iota + 1
	ConflictErrorType
	InternalServerErrorType
)
