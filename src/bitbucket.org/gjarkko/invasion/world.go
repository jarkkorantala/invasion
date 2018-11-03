package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"strings"
)

type World struct {
	Cities []*City
}

// Formats a World as string
func (world World) String() string {
	return fmt.Sprintf("World with %d cities: %v", len(world.Cities), world.Cities)
}

// Creates a World from a slice of Cities
// This is safer than calling the constructor directly as this checks for duplicate cities.
func CreateWorld(cities []*City) World {
	world := World{}

	// Sort cities for determinism
	cityMap := map[string]*City{}
	cityNames := []string{}
	for _, city := range cities {
		cityMap[city.Name] = city
		cityNames = append(cityNames, city.Name)
	}
	sort.Strings(cityNames)

	for _, cityName := range cityNames {
		world.AddCity(cityMap[cityName])
	}
	return world
}

// Add a City to the World
func (world *World) AddCity(city *City) {
	for _, existingCity := range world.Cities {
		if city == existingCity {
			log.Panicf("Panic: Invalid city to add - %s already exists in the world", city)
		}
	}
	world.Cities = append(world.Cities, city)
}

// Remove (destroy) a City from the World
func (world *World) RemoveCity(city *City) {
	for i, candidate := range world.Cities {
		if reflect.DeepEqual(city, candidate) {
			world.Cities = world.Cities[:i+copy(world.Cities[i:], world.Cities[i+1:])]
			return
		}
	}
	log.Panicf("Panic: Invalid city to delete - %s doesn't exist in the world", city)
}

// Pick a City in the given world at random
func (world World) RandomCity(randomizer Random) *City {
	return world.Cities[randomizer.Intn(len(world.Cities))]
}

// Serialize a World to string
func (world World) Serialize() string {
	cityStrings := []string{}
	for _, city := range world.Cities {
		cityStrings = append(cityStrings, city.Serialize())
	}

	// Ensure consistent ordering of cities for deterministic operation
	sort.Strings(cityStrings)
	serialized := ""
	for _, cityString := range cityStrings {
		serialized += fmt.Sprintf("%s\n", cityString)
	}
	return serialized

}

// Deserialize world from strings
func WorldFromString(serialized string) World {
	serialized = strings.Replace(serialized, "\r\n", "\n", -1)
	cityStrings := strings.Split(serialized, "\n")
	cityMap := map[string]*City{}
	neighborMap := map[*City]map[string]string{}

	// Parse serialized city strings, creating maps of cities and neighbour relationships
	for _, cityString := range cityStrings {

		// Ignore empty lines
		if cityString == "" {
			continue
		}

		city := CityFromString(cityString)
		cityMap[city.Name] = &city
		neighborMap[&city] = CityNeighborNamesFromString(cityString)
	}

	// Apply neighbor relationships
	for city, neighborNames := range neighborMap {
		for direction, neighborName := range neighborNames {
			city.SetNeighbor(direction, cityMap[neighborName])
		}
	}
	cities := []*City{}
	for _, city := range cityMap {
		cities = append(cities, city)
	}
	return CreateWorld(cities)
}

// Read World from a file
func WorldFromPath(path string) World {
	dat, err := ioutil.ReadFile(path)
	check(err)
	serialized := string(dat)
	world := WorldFromString(serialized)
	return world
}
