package main

import (
	"testing"
)

func TestCreateAlien(test *testing.T) {
	expected := CreateAlien("ET")
	actual := CreateAlien("ET")
	assertStringerEqual(test, expected, actual)
}

func TestCreateAlienInCity(test *testing.T) {
	sanFernando := CreateCity("San Fernando")
	alf := CreateAlienInCity("Alf", &sanFernando)
	assertStringerEqual(test, sanFernando, alf.City)
}

func TestMove(test *testing.T) {
	tokyo := CreateCity("Tokyo")
	godzilla := CreateAlien("Godzilla")
	godzilla.MoveToCity(&tokyo)
	assertStringerEqual(test, tokyo, godzilla.City)
}

func TestInvalidMovePanics(test *testing.T) {
	defer assertPanics(test)
	tokyo := CreateCity("Tokyo")
	godzilla := CreateAlien("Godzilla")
	godzilla.MoveToCity(&tokyo)
	godzilla.MoveToCity(&tokyo)
}

func TestMoveToRandomDirection(test *testing.T) {

	// Define a box of 4 cities
	sanFernando := CreateCity("San Fernando")
	burbank := CreateCity("Burbank")
	glendale := CreateCity("Glendale")
	calabasas := CreateCity("Calabasas")
	burbank.SetNeighbor(East, &sanFernando)
	burbank.SetNeighbor(South, &glendale)
	calabasas.SetNeighbor(North, &sanFernando)
	calabasas.SetNeighbor(East, &glendale)

	// Alien moves at known random seed
	randomizer := NewMockRandomizer()
	alf := CreateAlienInCity("Alf", &sanFernando)
	alf.MoveToRandomDirection(randomizer)
	assertStringerEqual(test, burbank, alf.City)
	alf.MoveToRandomDirection(randomizer)
	assertStringerEqual(test, glendale, alf.City)
	alf.MoveToRandomDirection(randomizer)
	assertStringerEqual(test, calabasas, alf.City)
}

func TestMoveToRandomDirectionWhenStuck(test *testing.T) {

	// Define an island
	tokyo := CreateCity("Tokyo")
	randomizer := NewMockRandomizer()
	godzilla := CreateAlienInCity("Godzilla", &tokyo)
	godzilla.MoveToRandomDirection(randomizer)
	godzilla.MoveToRandomDirection(randomizer)
	assertStringerEqual(test, tokyo, godzilla.City)
}

func TestMakeAlienName(test *testing.T) {
	assertStringEqual(test, "Alien 1", MakeAlienName(1))
	assertStringEqual(test, "Alien 51", MakeAlienName(51))
}
