package models_test

import (
	"testing"

	"github.com/santiagoh1997/weather-api/version-5/models"
)

func TestNewJob(t *testing.T) {
	tests := []struct {
		name    string
		city    string
		country string
	}{
		{"Bogotá, CO", "Bogotá", "CO"},
		{"Paris, FR", "Paris", "FR"},
		{"Amsterdam, NL", "Amsterdam", "NL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := models.NewJob(tt.city, tt.country)
			if j.City != tt.city {
				t.Errorf("NewJob Job.City = %v, want %v", j.City, tt.city)
			}
			if j.Country != tt.country {
				t.Errorf("NewJob Job.Country = %v, want %v", j.Country, tt.country)
			}
		})
	}
}
