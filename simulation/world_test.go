package main

import (
	"path"
	"reflect"
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

func TestRemoveCityByName(test *testing.T) {
	world := World{}
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world.AddCity(&moscow)
	world.AddCity(&vienna)
	world.RemoveCityByName("Moscow")
	expected := "World with 1 cities: [Vienna]"
	assertStringEqual(test, expected, world.String())
	world.RemoveCityByName("Vienna")
	expected = "World with 0 cities: []"
	assertStringEqual(test, expected, world.String())
}

func TestRemoveCityByNamePanicsOnNonExistent(test *testing.T) {
	defer assertPanics(test)
	world := World{}
	moscow := CreateCity("Moscow")
	world.AddCity(&moscow)
	world.RemoveCityByName("Vienna")
}

func TestRandomCity(test *testing.T) {
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	world := CreateWorld([]*City{&moscow, &vienna})

	randomizer := NewMockRandomizer()
	assertStringerEqual(test, moscow, world.RandomCity(randomizer))
	assertStringerEqual(test, moscow, world.RandomCity(randomizer))
	assertStringerEqual(test, moscow, world.RandomCity(randomizer))
}

func TestSerializeWorld(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	newYork := CreateCity("New York")
	scranton := CreateCity("Scranton")
	atlanticCity := CreateCity("Atlantic City")
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)
	atlanticCity.SetNeighbor(North, &philadelphia)
	newYork.SetNeighbor(West, &philadelphia)
	world := CreateWorld([]*City{&baltimore, &philadelphia, &newYork, &scranton, &atlanticCity})
	expected := "Atlantic City north=Philadelphia\n" +
		"Baltimore east=Philadelphia\n" +
		"New York west=Philadelphia\n" +
		"Philadelphia east=New York north=Scranton south=Atlantic City west=Baltimore\n" +
		"Scranton south=Philadelphia\n"
	assertStringEqual(test, expected, world.Serialize())
}

func TestWorldFromString(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	newYork := CreateCity("New York")
	scranton := CreateCity("Scranton")
	atlanticCity := CreateCity("Atlantic City")
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)
	atlanticCity.SetNeighbor(North, &philadelphia)
	newYork.SetNeighbor(West, &philadelphia)
	expected := CreateWorld([]*City{&baltimore, &philadelphia, &newYork, &scranton, &atlanticCity})
	actual := WorldFromString("Atlantic City north=Philadelphia\n" +
		"Baltimore east=Philadelphia\n" +
		"New York west=Philadelphia\n" +
		"Philadelphia east=New York north=Scranton south=Atlantic City west=Baltimore\n" +
		"Scranton south=Philadelphia\n")
	if !reflect.DeepEqual(expected, actual) {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

func TestWorldFromPath(test *testing.T) {
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	newYork := CreateCity("New York")
	scranton := CreateCity("Scranton")
	atlanticCity := CreateCity("Atlantic City")
	baltimore.SetNeighbor(East, &philadelphia)
	scranton.SetNeighbor(South, &philadelphia)
	atlanticCity.SetNeighbor(North, &philadelphia)
	newYork.SetNeighbor(West, &philadelphia)
	expected := CreateWorld([]*City{&baltimore, &philadelphia, &newYork, &scranton, &atlanticCity})

	fixturePath := path.Join("fixtures", "test_world.dat")
	actual := WorldFromPath(fixturePath)
	if !reflect.DeepEqual(expected, actual) {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}

}
