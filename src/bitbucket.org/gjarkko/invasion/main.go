package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const MaxIterations = 10000

// Run simulation for 10k iterations
func runSimulation(path string, alienCount int) {
	randomizer := NewRandomizer()
	simulation := SimulationFromPath(path, alienCount, randomizer)
	simulation.Simulate(MaxIterations)
	log.Printf("End of simulation. Remaining world:\n%s", simulation.World.Serialize())
}

// Parse user input and  launch simulation
func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Printf("Usage: %s [-n ALIENS] MAPFILE", os.Args[0])
		return
	}

	path := ""
	alienCount := 10
	for i := 0; i < len(argsWithoutProg); i++ {
		if argsWithoutProg[i] == "-n" {
			i, err := strconv.Atoi((argsWithoutProg[i+1]))
			check(err)
			i++
			continue
		} else {
			path = argsWithoutProg[i]
		}
	}
	log.Printf("Starting simulation for map file %s with %d aliens.", path, alienCount)
	runSimulation(path, alienCount)

}
