package main

import (
	"fmt"
	"log"
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
	for _, city := range cities {
		world.AddCity(city)
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
		if city == candidate {
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
