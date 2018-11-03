package main

import (
	"reflect"
	"testing"
)

func TestCreateCity(test *testing.T) {
	expected := CreateCity("London")
	actual := CreateCity("London")
	assertStringerEqual(test, expected, actual)
}

func TestCreateCityPanics(test *testing.T) {
	defer assertPanics(test)
	CreateCity("")
}

func TestNeighborMethods(test *testing.T) {
	washington := CreateCity("Washington DC")
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	scranton := CreateCity("Scranton")
	washington.SetNeighbor(East, &baltimore)
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)

	// Test east-west traverals
	expected := &philadelphia
	actual := washington.GetNeighbor(East).GetNeighbor(East)
	assertStringerEqual(test, expected, actual)
	expected = &washington
	actual = philadelphia.GetNeighbor(West).GetNeighbor(West)
	assertStringerEqual(test, expected, actual)

	// Test north-south traverals
	assertStringerEqual(test, &philadelphia, scranton.GetNeighbor(South))
	assertStringerEqual(test, &scranton, philadelphia.GetNeighbor(North))
}

func TestGetNeighborPanicsOninvalidDirection(test *testing.T) {
	defer assertPanics(test)
	washington := CreateCity("Washington DC")
	washington.GetNeighbor("up")
}

func TestSetNeighborPanicsOninvalidDirection(test *testing.T) {
	defer assertPanics(test)
	washington := CreateCity("Washington DC")
	baltimore := CreateCity("Baltimore")
	washington.SetNeighbor("Up", &baltimore)
}

func TestGetNeighborNilOnUnset(test *testing.T) {
	washington := CreateCity("Washington DC")
	assertNilCityReference(test, washington.GetNeighbor(North))
}

func TestSetNeighborPanicsOnConflictOnLocal(test *testing.T) {
	defer assertPanics(test)
	washington := CreateCity("Washington DC")
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	washington.SetNeighbor(East, &baltimore)
	baltimore.SetNeighbor(West, &philadelphia)
}
func TestSetNeighborPanicsOnConflictOnRemote(test *testing.T) {
	defer assertPanics(test)
	washington := CreateCity("Washington DC")
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	washington.Neighbors[East] = &baltimore
	philadelphia.SetNeighbor(West, &washington)
}

func TestDestroy(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	scranton := CreateCity("Scranton")
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)

	philadelphia.Destroy()
	assertNilCityReference(test, baltimore.GetNeighbor(East))
	assertNilCityReference(test, scranton.GetNeighbor(South))
}

func TestSerializeOrphanCity(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	assertStringEqual(test, "Baltimore", baltimore.Serialize())
}

func TestSerializeCityWithNeighbors(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	newYork := CreateCity("New York")
	scranton := CreateCity("Scranton")
	atlanticCity := CreateCity("Atlantic City")
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)
	atlanticCity.SetNeighbor(North, &philadelphia)
	newYork.SetNeighbor(West, &philadelphia)
	assertStringEqual(test, "New York west=Philadelphia", newYork.Serialize())
	assertStringEqual(test, "Baltimore east=Philadelphia", baltimore.Serialize())
	assertStringEqual(test, "Philadelphia east=New York north=Scranton south=Atlantic City west=Baltimore", philadelphia.Serialize())
	assertStringEqual(test, "Scranton south=Philadelphia", scranton.Serialize())
	assertStringEqual(test, "Atlantic City north=Philadelphia", atlanticCity.Serialize())
}

func TestCityFromString(test *testing.T) {
	city := CityFromString("New York")
	assertStringEqual(test, "New York", city.Name)
	city = CityFromString("Philadelphia east=New York north=Scranton south=Atlantic City west=Baltimore")
	assertStringEqual(test, "Philadelphia", city.Name)
}

func TestCityNeighborNamesFromString(test *testing.T) {
	neighborNames := CityNeighborNamesFromString("Baltimore")
	expected := map[string]string{}
	if !reflect.DeepEqual(expected, neighborNames) {
		test.Errorf("expected: %+v, actual: %+v", expected, neighborNames)
	}
	neighborNames = CityNeighborNamesFromString("Philadelphia east=New York north=Scranton south=Atlantic City west=Baltimore")
	expected = map[string]string{"east": "New York", "north": "Scranton", "south": "Atlantic City", "west": "Baltimore"}
	if !reflect.DeepEqual(expected, neighborNames) {
		test.Errorf("expected: %+v, actual: %+v", expected, neighborNames)
	}
}
