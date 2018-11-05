package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

const North = "north"
const East = "east"
const South = "south"
const West = "west"

var Directions = [4]string{North, East, South, West}
var opposites = map[string]string{
	North: South,
	East:  West,
	South: North,
	West:  East,
}

// Defines a city with a name and optional neighbor cities in cardinal directions.
type City struct {
	Name      string
	Neighbors map[string]*City
}

// Formats a City as string
func (city City) String() string {
	return city.Name
}

// Create named city
func CreateCity(name string) City {
	if name == "" {
		panic("No name given for City.")
	}
	neighbors := make(map[string]*City)
	city := City{name, neighbors}
	debugLog(fmt.Sprintf("City %s was created.", city))
	return city
}

// Sets a neighbor to a given direction (two-directional)
func (city City) SetNeighbor(direction string, neighbor *City) {
	// Validate direction
	if !isDirection(direction) {
		panic("Invalid direction \"" + direction + "\"")
	}

	// Check for conflicts previously set
	if old, ok := city.Neighbors[direction]; ok {
		if !reflect.DeepEqual(old, neighbor) {
			log.Panicf(fmt.Sprintf("Panic: Invalid neighbor definition for %s: %s=%s. "+
				"Already set to %s.", city, direction, neighbor, old))
		}
	}

	// Check for conflicting setup on neighbor
	opposite := opposites[direction]
	if old, ok := neighbor.Neighbors[opposite]; ok {
		if !reflect.DeepEqual(old, &city) {
			log.Panicf(fmt.Sprintf("Panic: Invalid neighbor definition for %s: %s=%s. Already set to %s.",
				neighbor, opposite, city, old))
		}
	}
	city.Neighbors[direction] = neighbor
	neighbor.Neighbors[opposite] = &city
	debugLog(fmt.Sprintf("City %s has a road to neighbour %s to the %s.", city, neighbor, direction))
}

// Gets a neighbor for a given direction
func (city City) GetNeighbor(direction string) *City {
	if !isDirection(direction) {
		panic("Invalid direction \"" + direction + "\"")
	}
	if neighbor, ok := city.Neighbors[direction]; ok {
		return neighbor
	}
	return nil
}

// Marks a city destroyed, removing it from neighbors
func (city City) Destroy() {
	for direction, neighbor := range city.Neighbors {
		if neighbor == nil {
			continue
		}
		opposite := opposites[direction]
		neighbor.Neighbors[opposite] = nil
		debugLog(fmt.Sprintf("City %s is no longer a neighbor of %s in %s.\n", city, neighbor, opposite))
	}
	log.Printf("City %s was destroyed.\n", city)
}

// Tests that direction is supported
func isDirection(direction string) bool {
	for _, validDirection := range Directions {
		if direction == validDirection {
			return true
		}
	}
	return false
}

// Serialize a City to string
func (city City) Serialize() string {

	serialized := city.Name

	// Find directions that have a neighbor
	available := []string{}
	for direction, _ := range city.Neighbors {
		available = append(available, direction)
	}

	// Ensure consistent ordering of directions for deterministic operation
	sort.Strings(available)

	for _, direction := range available {
		neighbor := city.Neighbors[direction]
		if neighbor == nil {
			continue
		}
		serialized += fmt.Sprintf(" %s=%s", direction, neighbor)
	}
	return serialized
}

// Deserialize a City from string (without neighbors)
func CityFromString(serialized string) City {

	// Find positions of cardinal directories ("direction=")
	re := regexp.MustCompile("( [\\w]+=)")
	positions := re.FindAllStringIndex(serialized, -1)
	cityName := serialized
	if len(positions) > 0 {
		cityName = serialized[0:positions[0][0]]
	}
	city := CreateCity(cityName)
	return city
}

// Deserialize names of neighbors from string
func CityNeighborNamesFromString(serialized string) map[string]string {

	// Find positions of cardinal directories ("direction=")
	re := regexp.MustCompile("( [\\w]+=)")
	positions := re.FindAllStringIndex(serialized, -1)

	// From parts separated by these positions, parse direction and city name
	neighborNames := map[string]string{}
	for i, indices := range positions {
		start := indices[0] + 1
		end := len(serialized)
		if i+1 < len(positions) {
			end = positions[i+1][0]
		}
		parts := strings.Split(serialized[start:end], "=")
		neighborNames[parts[0]] = parts[1]
	}
	return neighborNames
}
