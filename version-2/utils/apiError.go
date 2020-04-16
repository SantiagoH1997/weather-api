package utils

import (
	"net/http"
)

const internalServerErrorMessage = "Internal server error"

// APIError is the error returned to the user
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// NewInternalServerError returns an APIError
// with a 500 status code and a message
func NewInternalServerError() *APIError {
	return &APIError{
		http.StatusInternalServerError,
		internalServerErrorMessage,
	}
}

// NewNotFound returns an APIError with a 404 status code and a message
func NewNotFound(message string) *APIError {
	return &APIError{
		http.StatusNotFound,
		message,
	}
}

// NewBadRequest returns an APIError with a 400 status code and a message
func NewBadRequest(message string) *APIError {
	return &APIError{
		http.StatusBadRequest,
		message,
	}
}

// NewApiError returns a custom APIError
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		statusCode,
		message,
	}
}
