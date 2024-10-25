package constants

type ErrorCode int

const (
	ErrorCodeValidation ErrorCode = iota + 1
	ErrorCodeConflict
	ErrorCodeInternalServer
)

type Case string

const (
	CaseInvalidUUIDLength Case = "invalid_uuid_length"
)
