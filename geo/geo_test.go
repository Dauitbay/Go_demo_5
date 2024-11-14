package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	//Act
	got,err := geo.GetMyLocation(city)
	//Assert
	if err != nil {
		t.Error("Error of getting city")
	}
	if got.City != expected.City{
		t.Errorf("Expected %v, but got %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Moscow"
	_,err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity{
		t.Errorf("Ecpected %v, but got %v",geo.ErrNot200, err)
	}
}