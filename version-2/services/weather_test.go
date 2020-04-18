package services_test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/santiagoh1997/weather-api/version-2/services"
	"github.com/santiagoh1997/weather-api/version-2/testdata"
)

func TestNewWeatherService(t *testing.T) {
	db, teardown, err := setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, nil)
	if ws.Database == nil {
		t.Errorf("NewWeatherService want %v, got %v", db, nil)
	}
}

func TestGetByLocation(t *testing.T) {
	db, teardown, err := setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	ws := services.NewWeatherService(db, nil)

	t.Run("Success", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"Bogotá", "Bogotá, CO"},
			{"Paris", "Paris, FR"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w, err := ws.GetByLocation(tt.input)
				if err != nil {
					t.Errorf("GetByLocation err = %v, want %v", err, nil)
				}
				if w == nil {
					t.Errorf("GetByLocation = %v, want *Weather", w)
				}
			})
		}
	})

	t.Run("Error", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
		}{
			{"Buenos Aires", "Buenos Aires, AR"},
			{"Lima", "Lima, PE"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w, err := ws.GetByLocation(tt.input)
				if err == nil {
					t.Errorf("GetByLocation err = %v, want APIError", err)
				}
				if w != nil {
					t.Errorf("GetByLocation = %v, want %v", w, nil)
				}
			})
		}
	})

}

func TestSave(t *testing.T) {
	db, teardown, err := setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, nil)

	t.Run("Success", func(t *testing.T) {
		testWeather := testdata.SampleWeather
		id := primitive.NewObjectID()
		testWeather.ID = id
		res, err := ws.Save(&testWeather)
		if err != nil {
			t.Errorf("Save error = %v, want %v", err, nil)
		}
		if res.InsertedID != id {
			t.Errorf("Save InsertedID = %v, want %v", res.InsertedID, id)
		}
	})
}

func TestUpdate(t *testing.T) {
	db, teardown, err := setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, nil)

	t.Run("Success", func(t *testing.T) {
		testWeather := testdata.TestWeathers[0]
		testWeather.Temperature = "20 °C"
		res, err := ws.Update(&testWeather)
		if err != nil {
			t.Errorf("Update error = %v, want %v", err, nil)
		}
		if res.ModifiedCount != 1 {
			t.Errorf("Update ModifiedCount = %v, want %v", res.ModifiedCount, 1)
		}
	})
}
