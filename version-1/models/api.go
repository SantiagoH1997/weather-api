package models

// APIError is the error returned to the user
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
