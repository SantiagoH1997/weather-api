package db_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-3/db"
)

var (
	mongoURI string
	dbName   string
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	beego.TestBeegoInit(filepath.Dir(pwd))
	mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoTestHost"), beego.AppConfig.String("mongoPort"))
	dbName = beego.AppConfig.String("mongoTestDBName")
}

func TestOpen(t *testing.T) {
	tests := []struct {
		name string
		URI  string
		DB   string
	}{
		{"Success", mongoURI, dbName},
		{"Wrong DB name", mongoURI, ""},
		{"Wrong URI", "", dbName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, close, err := db.Open(tt.URI, tt.DB)
			if err != nil {
				if tt.URI != "" && tt.DB != "" {
					t.Errorf("Open err = %v, want %v", err, nil)
				}
			} else {
				defer close(context.Background())
			}
		})
	}
}
