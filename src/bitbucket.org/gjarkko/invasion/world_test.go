package main

import (
	"testing"
)

func TestCreateWorld(test *testing.T) {
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world := CreateWorld([]*City{&moscow, &vienna})
	expected := "World with 2 cities: [Moscow Vienna]"
	assertStringEqual(test, expected, world.String())
}

func TestAddCity(test *testing.T) {
	world := World{}
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world.AddCity(&moscow)
	world.AddCity(&vienna)
	expected := "World with 2 cities: [Moscow Vienna]"
	assertStringEqual(test, expected, world.String())
}

func TestAddCityPanicsOnDuplicate(test *testing.T) {
	defer assertPanics(test)
	world := World{}
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world.AddCity(&moscow)
	world.AddCity(&vienna)
	world.AddCity(&moscow)
	expected := "World with 2 cities: [Moscow Vienna]"
	assertStringEqual(test, expected, world.String())
}

func TestRemoveCity(test *testing.T) {
	world := World{}
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world.AddCity(&moscow)
	world.AddCity(&vienna)
	world.RemoveCity(&moscow)
	expected := "World with 1 cities: [Vienna]"
	assertStringEqual(test, expected, world.String())
	world.RemoveCity(&vienna)
	expected = "World with 0 cities: []"
	assertStringEqual(test, expected, world.String())
}

func TestRemoveCityPanicsOnNonExistent(test *testing.T) {
	defer assertPanics(test)
	world := World{}
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world.AddCity(&moscow)
	world.RemoveCity(&vienna)
}

func TestRandomCity(test *testing.T) {
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world := CreateWorld([]*City{&moscow, &vienna})

	randomizer := NewMockRandomizer()
	assertStringerEqual(test, vienna, world.RandomCity(randomizer))
	assertStringerEqual(test, vienna, world.RandomCity(randomizer))
	assertStringerEqual(test, vienna, world.RandomCity(randomizer))
}
