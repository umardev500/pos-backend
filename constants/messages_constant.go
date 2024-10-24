package constants

type MessageText string

const (
	// Common
	MessageCommonValidationFailed    MessageText = "validation failed"
	MessageCommonInternalServerError MessageText = "internal server error"

	// User service
	MessageUserCreateSuccess MessageText = "user created successfully"
	MessageUserCreateFailed  MessageText = "user creation failed"
)
