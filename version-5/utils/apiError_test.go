package utils_test

import (
	"net/http"
	"testing"

	"github.com/santiagoh1997/weather-api/version-5/utils"
)

const (
	messageInternalServerError = "Internal server error"
	messageNotFound            = "Not found"
	messageBadRequest          = "Bad request"
	apiErrorStatusCode         = 507
	messageApiError            = "Random message"
)

func TestNewInternalServerError(t *testing.T) {
	err := utils.NewInternalServerError()
	if err.StatusCode != http.StatusInternalServerError {
		t.Errorf("NewInternalServerError err.StatusCode = %v, want %v", err.StatusCode, http.StatusInternalServerError)
	}
	if err.Message != messageInternalServerError {
		t.Errorf("NewInternalServerError err.Message = %v, want %v", err.Message, messageInternalServerError)
	}
}

func TestNewNotFound(t *testing.T) {
	err := utils.NewNotFound(messageNotFound)
	if err.StatusCode != http.StatusNotFound {
		t.Errorf("NewNotFound err.StatusCode = %v, want %v", err.StatusCode, http.StatusNotFound)
	}
	if err.Message != messageNotFound {
		t.Errorf("NewNotFound err.Message = %v, want %v", err.Message, messageNotFound)
	}
}

func TestNewBadRequest(t *testing.T) {
	err := utils.NewBadRequest(messageBadRequest)
	if err.StatusCode != http.StatusBadRequest {
		t.Errorf("NewBadRequest err.StatusCode = %v, want %v", err.StatusCode, http.StatusBadRequest)
	}
	if err.Message != messageBadRequest {
		t.Errorf("NewBadRequest err.Message = %v, want %v", err.Message, messageBadRequest)
	}
}

func TestNewAPIError(t *testing.T) {
	err := utils.NewAPIError(apiErrorStatusCode, messageApiError)
	if err.StatusCode != apiErrorStatusCode {
		t.Errorf("NewAPIError err.StatusCode = %v, want %v", err.StatusCode, apiErrorStatusCode)
	}
	if err.Message != messageApiError {
		t.Errorf("NewAPIError err.Message = %v, want %v", err.Message, messageApiError)
	}
}
