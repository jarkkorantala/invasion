package main

import (
	"fmt"
	"log"
	"sort"
)

type Alien struct {
	Name string
	City *City
}

// Formats an Alien as string
func (alien Alien) String() string {
	return fmt.Sprintf("%s", alien.Name)
}

// Creates a named Alien
func CreateAlien(name string) Alien {
	alien := Alien{name, nil}
	debugLog(fmt.Sprintf("Alien %s created.", alien))
	return alien
}

// Creates a named Alien
func CreateAlienInCity(name string, city *City) Alien {
	alien := Alien{name, city}
	debugLog(fmt.Sprintf("Alien %s created.", alien))
	return alien
}

// Moves the given alien to the given city
func (alien *Alien) MoveToCity(city *City) {
	if alien.City == city {
		log.Panicf("Panic: Invalid move for alien %s; already in %s.", alien, city)
	}
	alien.City = city
	debugLog(fmt.Sprintf("Alien %s moved to %s.", alien, city))
}

// Moves the given alien to a random available direction
func (alien *Alien) MoveToRandomDirection(randomizer Random) {

	// Find directions that have a neighbor
	available := []string{}
	for direction, city := range alien.City.Neighbors {
		if city != nil {
			available = append(available, direction)
		}
	}

	if len(available) == 0 {
		debugLog(fmt.Sprintf("Alien %s is stuck.", alien))
		return
	} else {

		// Ensure consistent ordering of directions for deterministic operation
		sort.Strings(available)

		direction := available[randomizer.Intn(len(available))]
		target := alien.City.Neighbors[direction]
		alien.MoveToCity(target)
	}
}

// Creates an alien name for the given seed
func MakeAlienName(i int) string {
	return fmt.Sprintf("Alien %d", i)
}
