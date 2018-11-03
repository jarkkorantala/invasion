package main

import "testing"

func createMockSimulation(alienCount int) Simulation {
	moscow := CreateCity("Moscow")
	vienna := CreateCity("Vienna")
	vienna.SetNeighbor(East, &moscow)
	randomizer := NewMockRandomizer()
	simulation := CreateSimulation([]*City{&moscow, &vienna}, alienCount, randomizer)

	return simulation
}

func TestCreateSimulation(test *testing.T) {
	simulation := createMockSimulation(12)
	expected := "Simulation with World with 2 cities: [Moscow Vienna] and 12 aliens"
	assertStringEqual(test, expected, simulation.String())
	for _, alien := range simulation.Aliens {
		assertStringEqual(test, "Moscow", alien.City.String())
	}
}

func TestRemoveAlien(test *testing.T) {
	simulation := createMockSimulation(12)

	expected := "Simulation with World with 2 cities: [Moscow Vienna] and 12 aliens"
	assertStringEqual(test, expected, simulation.String())
	for _, alien := range simulation.Aliens {
		assertStringEqual(test, "Moscow", alien.City.String())
	}
}

func TestRemoveAlienPanicsOnUnknownAlien(test *testing.T) {
	simulation := createMockSimulation(12)
	defer assertPanics(test)
	alf := CreateAlien("Alf")
	simulation.RemoveAlien(&alf)
}

func TestAdvanceAliens(test *testing.T) {
	simulation := createMockSimulation(12)
	assertIntEqual(test, 0, simulation.Iteration)
	for _, alien := range simulation.Aliens {
		assertStringEqual(test, "Moscow", alien.City.String())
	}
	simulation.AdvanceAliens()
	assertIntEqual(test, 12, len(simulation.Aliens))
	assertIntEqual(test, 2, len(simulation.World.Cities))
	for _, alien := range simulation.Aliens {
		assertStringEqual(test, "Vienna", alien.City.String())
	}
}

func TestAdvance(test *testing.T) {
	simulation := createMockSimulation(12)

	// Add one odd alien that's lost from  others
	alf := CreateAlienInCity("Alf", simulation.World.Cities[1])

	simulation.AddAlien(&alf)
	assertIntEqual(test, 0, simulation.Iteration)
	for i := 0; i < 12; i++ {
		assertStringEqual(test, "Moscow", simulation.Aliens[i].City.String())
	}
	assertStringEqual(test, "Vienna", simulation.Aliens[12].City.String())
	simulation.Advance()
	assertIntEqual(test, 1, simulation.Iteration)
	assertIntEqual(test, 1, len(simulation.Aliens))
	assertIntEqual(test, 1, len(simulation.World.Cities))
	assertStringEqual(test, "Moscow", simulation.Aliens[0].City.String())
}

func TestSimulateToMaxIterations(test *testing.T) {
	simulation := createMockSimulation(1)
	simulation.Simulate(4)
	assertIntEqual(test, 4, simulation.Iteration)
	assertIntEqual(test, 1, len(simulation.Aliens))
	assertIntEqual(test, 2, len(simulation.World.Cities))

}

func TestSimulateToArmageddon(test *testing.T) {
	simulation := createMockSimulation(2)
	cthulu := CreateAlienInCity("Cthulu", simulation.World.Cities[1])
	shubby := CreateAlienInCity("Shub-Niggurath", simulation.World.Cities[1])
	simulation.AddAlien(&cthulu)
	simulation.AddAlien(&shubby)
	simulation.Simulate(4)
	assertIntEqual(test, 1, simulation.Iteration)
	assertIntEqual(test, 0, len(simulation.Aliens))
	assertIntEqual(test, 0, len(simulation.World.Cities))
}
