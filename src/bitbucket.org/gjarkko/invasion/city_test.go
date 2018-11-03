package main

import (
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
