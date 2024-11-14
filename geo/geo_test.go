package geo_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	//Act
	got, err := geo.GetMyLocation(city)
	//Assert
	if err != nil {
		t.Error("Error of getting city")
	}
	if got.City != expected.City {
		t.Errorf("Expected %v, but got %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Moscow"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ecpected %v, but got %v", geo.ErrNot200, err)
	}
}

func TestGetWeather(t *testing.T) {
	expected := "Nukus"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Got error %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v , but got %v", expected, result)
	}
}
 
var testCases = []struct {
	name string
	format int
}{
	{ name: "Big format", format: 147 },
	{ name: "0 format", format: 0 },
	{ name: "Minus format", format: -1},
}
func TestGetWeatherWrongFormat(t *testing.T) {
	for _, testC := range testCases {
		t.Run(testC.name, func(t *testing.T){
			expected := "Nukus"
			geoData := geo.GeoData{
				City: expected,
			}
			format := 125
			_, err := weather.GetWeather(geoData, format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Expected %v , but got %v", weather.ErrWrongFormat, err)
			}
		})
	}
	
}
