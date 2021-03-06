package services_test

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-5/logger"
	"github.com/santiagoh1997/weather-api/version-5/services"
	"github.com/santiagoh1997/weather-api/version-5/testdata"
	"github.com/santiagoh1997/weather-api/version-5/testutils"
	"github.com/santiagoh1997/weather-api/version-5/utils"
)

func TestNewWeatherService(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, nil, nil)
	if ws.Database == nil {
		t.Errorf("NewWeatherService.Database want %v, got %v", db, nil)
	}
}

func TestGetByLocation(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	ws := services.NewWeatherService(db, nil, nil)

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
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, nil, nil)

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
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := services.NewWeatherService(db, logger.NewLogger(), nil)

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

func TestFetchWeather(t *testing.T) {
	expectedError := &utils.APIError{StatusCode: 404, Message: ""}
	tests := []struct {
		name string
		URL  string
		err  *utils.APIError
		city string
	}{
		{"Paris", fmt.Sprintf(testutils.APIURL, "Paris", "FR"), nil, "Paris"},
		{"Amsterdam", fmt.Sprintf(testutils.APIURL, "Amsterdam", "NL"), nil, "Amsterdam"},
		{"Rotterdam", fmt.Sprintf(testutils.APIURL, "Rotterdam", "NL"), nil, "Rotterdam"},
		{"Abc123", fmt.Sprintf(testutils.APIURL, "Abc123", "NL"), expectedError, ""},
		{"Def456", fmt.Sprintf(testutils.APIURL, "Def456", "NZ"), expectedError, ""},
	}
	ws := services.NewWeatherService(nil, nil, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := ws.FetchWeather(tt.URL)
			if err != nil && tt.err != nil {
				if err.StatusCode != tt.err.StatusCode {
					t.Errorf("FetchWeather err.StatusCode = %v, want %v", err.StatusCode, tt.err.StatusCode)
				}
			}
			if w != nil {
				if tt.city == "" {
					t.Errorf("FetchWeather Weather.City = %v, want %q", w.City, tt.city)
				}
				got := w.City
				if got != tt.city {
					t.Errorf("FetchWeather got = %v, want %v", got, tt.city)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	tests := []struct {
		name    string
		city    string
		country string
		want    string
	}{
		{"Bogotá", "Bogotá", "CO", "Bogotá, CO"},
		{"Paris", "Paris", "FR", "Paris, FR"},
		{"São Paulo", "São Paulo", "BR", "São Paulo, BR"},
	}
	ws := services.NewWeatherService(db, nil, nil)
	services.APIURL = testutils.APIURL

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := ws.Get(tt.city, tt.country)
			if err != nil {
				t.Fatalf("Get err = %v, want %v", err, nil)
			}
			got := w.LocationName
			if got != tt.want {
				t.Errorf("Get got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetFromJSON(t *testing.T) {
	beego.AppConfig.Set("dataSource", "json")
	defer beego.AppConfig.Set("dataSource", "DB")

	tests := []struct {
		name    string
		city    string
		country string
		want    string
	}{
		{"Bogotá", "Bogotá", "CO", "Bogotá, CO"},
		{"Paris", "Paris", "FR", "Paris, FR"},
	}
	ws := services.NewWeatherService(nil, nil, nil)
	services.APIURL = testutils.APIURL

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := ws.Get(tt.city, tt.country)
			if err != nil {
				t.Fatalf("Get err = %v, want %v", err, nil)
			}
			got := w.LocationName
			if got != tt.want {
				t.Errorf("Get got = %q, want %q", got, tt.want)
			}
		})
	}
}
