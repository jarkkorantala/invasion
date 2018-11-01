package main

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
	Destroyed bool
}

// Formats a City as string
func (city City) String() string {
	return city.Name
}

// Create named city
func CreateCity(name string) City {
	neighbors := make(map[string]*City)
	return City{name, neighbors, false}
}

// Sets a neighbor to a given direction (two-directional)
func (city City) SetNeighbor(direction string, neighbor *City) {
	if !isDirection(direction) {
		panic("Invalid direction \"" + direction + "\"")
	}
	city.Neighbors[direction] = neighbor
	neighbor.Neighbors[opposites[direction]] = &city
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

// Marks a city destroyed
func (city City) Destroy() {
	city.Destroyed = true
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
