package utils

import "strings"

// Request is the struct for the user's request to the API
type Request struct {
	City    string
	Country string
}

// ValidateRequest validates the fields of a request
func (r *Request) ValidateRequest() *APIError {
	if strings.TrimSpace(r.City) == "" {
		if strings.TrimSpace(r.Country) == "" {
			return NewBadRequest("Missing params: city, country")
		}
		return NewBadRequest("Missing param: city")
	}
	if strings.TrimSpace(r.Country) == "" {
		return NewBadRequest("Missing param: country")
	}
	return nil
}
