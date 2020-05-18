// Package doc Weather API.
//
// API to request weather reports
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	License: MIT http://opensource.org/licenses/MIT
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package doc

// swagger:parameters getWeather
type getWeatherParameters struct {
	// City name
	// Example: Bogotá
	// in: path
	// required: true
	City string `json:"city"`
	// Two letter country code
	// Example: CO
	// in: path
	// required: true
	Country string `json:"country"`
}

// swagger:parameters scheduleJob
type scheduleJobParameters struct {
	// City name
	// Example: Bogotá
	// in: body
	// required: true
	City string `json:"city"`
	// Country name
	// Two letter country code
	// Example: CO
	// in: body
	// required: true
	Country string `json:"country"`
}

// notFoundError
// swagger:response notFoundError
// in: body
type notFoundError struct {
	// 404
	StatusCode int `json:"status_code"`
	// City not found
	Message string `json:"message"`
}

// badRequestError
// swagger:response badRequestError
// in: body
type badRequestError struct {
	// 400
	StatusCode int `json:"status_code"`
	// A list of missing fields
	Message string `json:"message"`
}

// jobScheduledResponse
// swagger:response jobScheduledResponse
// in: body
type jobScheduledResponse struct {
	// 202
	StatusCode int `json:"status_code"`
	// "Job scheduled"
	Message string `json:"message"`
}
