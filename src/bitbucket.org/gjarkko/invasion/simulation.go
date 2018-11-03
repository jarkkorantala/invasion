package main

import (
	"fmt"
	"log"
)

type Simulation struct {
	World      *World
	Aliens     []*Alien
	Iteration  int
	randomizer Random
}

// Formats a City as string
func (simulation Simulation) String() string {
	return fmt.Sprintf("Simulation with %s and %v aliens", simulation.World, len(simulation.Aliens))
}

// Create a number of aliens for a world
func createAliens(world World, alienCount int, randomizer Random) []*Alien {
	var aliens []*Alien
	for i := 0; i < alienCount; i++ {
		alienName := MakeAlienName(i)
		originCity := world.RandomCity(randomizer)
		alien := CreateAlienInCity(alienName, originCity)
		aliens = append(aliens, &alien)
	}
	return aliens
}

// Creates a simulation for the given cities for the given number of aliens
func CreateSimulation(cities []*City, alienCount int, randomizer Random) Simulation {
	world := CreateWorld(cities)
	aliens := createAliens(world, alienCount, randomizer)
	simulation := Simulation{&world, aliens, 0, randomizer}
	log.Printf("Created %s\n", simulation)
	return simulation
}

// Creates simulation for the given map  path and for the given number of aliens
func SimulationFromPath(path string, alienCount int, randomizer Random) Simulation {
	world := WorldFromPath(path)
	aliens := createAliens(world, alienCount, randomizer)
	simulation := Simulation{&world, aliens, 0, randomizer}
	log.Printf("Created %s\n", simulation)
	return simulation
}

// Advances the aliens to the next iteration
func (simulation *Simulation) AdvanceAliens() {
	for _, alien := range simulation.Aliens {
		alien.MoveToRandomDirection(simulation.randomizer)
	}
}

// Add an Alien to the Simulation
func (simulation *Simulation) AddAlien(alien *Alien) {
	simulation.Aliens = append(simulation.Aliens, alien)
}

// Remove (destroy) an Alien from the Simulation
func (simulation *Simulation) RemoveAlien(alien *Alien) {
	for i, candidate := range simulation.Aliens {
		if alien == candidate {
			simulation.Aliens = simulation.Aliens[:i+copy(simulation.Aliens[i:], simulation.Aliens[i+1:])]
			return
		}
	}
	log.Panicf("Panic: Invalid alien to delete - %s doesn't exist in the simulation", alien)
}

// Checks the world for multiple aliens in the same city, removes aliens and city when found
func (simulation *Simulation) ResolveConflicts() {

	// Group aliens by city
	presenceMap := make(map[*City][]*Alien)
	for _, alien := range simulation.Aliens {
		presenceMap[alien.City] = append(presenceMap[alien.City], alien)
	}

	// Destory cities that have multiple aliens
	for city, aliens := range presenceMap {
		if len(aliens) <= 1 {
			continue
		}
		log.Printf("Whoa! Multiple aliens in %s: %s.", city, aliens)
		simulation.World.RemoveCity(city)
		for _, alien := range aliens {
			simulation.RemoveAlien(alien)
			log.Printf("%s was killed.", alien)
		}
	}
}

// Advances the Simulation to the next iteration by moving aliens and resolving conflicts
func (simulation *Simulation) Advance() {
	simulation.AdvanceAliens()
	simulation.ResolveConflicts()
	simulation.Iteration += 1
}

// Run simulation until no more aliens left or max iterations reached
func (simulation *Simulation) Simulate(maxIteration int) {
	for {
		simulation.Advance()
		if simulation.Iteration == maxIteration {
			log.Printf("Maximum number of iterations reached (%d)", maxIteration)
			break
		}
		if len(simulation.World.Cities) == 0 {
			log.Printf("No more cities left in the world.")
			break
		}
	}
}
