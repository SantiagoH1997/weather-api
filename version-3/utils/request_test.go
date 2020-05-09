package utils_test

import (
	"testing"

	"github.com/santiagoh1997/weather-api/version-3/utils"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		request *utils.Request
		want    string
	}{
		{"Missing city", &utils.Request{City: "", Country: "Barbados"}, "Missing param: city"},
		{"Missing country", &utils.Request{City: "Sofia", Country: ""}, "Missing param: country"},
		{"Missing city and country", &utils.Request{City: "", Country: ""}, "Missing params: city, country"},
		{"Correct input", &utils.Request{City: "Kosovo", Country: "BA"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.ValidateRequest()
			if err != nil {
				if tt.request.City != "" && tt.request.Country != "" {
					t.Errorf("ValidateRequest err = %v, want %v", err, nil)
				}
				got := err.Message
				if got != tt.want {
					t.Errorf("ValidateRequest err = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
