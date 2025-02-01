package errors

import "fmt"

// ErrorType represents the type of error that occurred
type ErrorType string

const (
	ValidationError ErrorType = "VALIDATION_ERROR"
	RuntimeError    ErrorType = "RUNTIME_ERROR"
	SystemError     ErrorType = "SYSTEM_ERROR"
)

// AppError represents a structured error response
type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Detail  string    `json:"detail"`
	Code    int       `json:"code"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Type, e.Message, e.Detail)
}

// New creates a new AppError
func New(errorType ErrorType, message string, detail string, code int) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
		Detail:  detail,
		Code:    code,
	}
}

// Common error constructors
func NewValidationError(message string, detail string) *AppError {
	return New(ValidationError, message, detail, 400)
}

func NewRuntimeError(message string, detail string) *AppError {
	return New(RuntimeError, message, detail, 500)
}

func NewSystemError(message string, detail string) *AppError {
	return New(SystemError, message, detail, 500)
}
