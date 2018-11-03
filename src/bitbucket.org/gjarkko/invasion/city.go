package main

import (
	"fmt"
	"log"
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
	log.Printf("City %s was created.", city)
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
		if old != neighbor {
			log.Panicf(fmt.Sprintf("Panic: Invalid neighbor definition for %s: %s=%s. "+
				"Already set to %s.", city, direction, neighbor, old))
		}
	}

	// Check for conflicting setup on neighbor
	opposite := opposites[direction]
	if old, ok := neighbor.Neighbors[opposite]; ok {
		if old != &city {
			log.Panicf(fmt.Sprintf("Panic: Invalid neighbor definition for %s: %s=%s. Already set to %s.",
				neighbor, opposite, city, old))
		}
	}
	city.Neighbors[direction] = neighbor
	neighbor.Neighbors[opposite] = &city
	log.Printf("City %s has a road to neighbour %s to the %s.", city, neighbor, direction)
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
		opposite := opposites[direction]
		neighbor.Neighbors[opposite] = nil
		log.Printf("City %s is no longer a neighbor of %s in %s.\n", city, neighbor, opposite)
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
