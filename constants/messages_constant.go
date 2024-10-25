package constants

type MessageText string

const (
	// Common
	MessageCommonValidationFailed    MessageText = "validation failed"
	MessageCommonInternalServerError MessageText = "internal server error"
	MessageCommonUUIDInvalidLength   MessageText = "invalid uuid length"
	MessageCommonBadRequest          MessageText = "bad request"

	// User service
	MessageUserCreateSuccess MessageText = "user created successfully"
	MessageUserCreateFailed  MessageText = "user creation failed"
	MessageUserDeleteSuccess MessageText = "user deleted successfully"
	MessageUserDeleteFailed  MessageText = "user deletion failed"
)
