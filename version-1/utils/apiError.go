package utils

import (
	"net/http"

	"github.com/santiagoh1997/weather-api/version-1/logger"
)

const internalServerErrorMessage = "Internal server error"

// APIError is the error returned to the user
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// NewInternalServerError logs an error and
// returns an APIError with a 500 status code and a message
func NewInternalServerError(logError string) *APIError {
	logger.Log.Errorf(logError)
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

// NewApiError returns a custom APIError
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		statusCode,
		message,
	}
}
