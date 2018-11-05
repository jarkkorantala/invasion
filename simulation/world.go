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
			city.Destroy()
			return
		}
	}
	log.Panicf("Panic: Invalid city to delete - %s doesn't exist in the world", city)
}

// Remove (destroy) a City from the World by it's name
func (world *World) RemoveCityByName(cityName string) {
	for i, city := range world.Cities {
		if city.Name == cityName {
			world.Cities = world.Cities[:i+copy(world.Cities[i:], world.Cities[i+1:])]
			city.Destroy()
			return
		}
	}
	log.Panicf("Panic: Invalid city to delete - %s doesn't exist in the world", cityName)
}

// Pick a City in the given world at random
func (world World) RandomCity(randomizer Random) *City {
	return world.Cities[randomizer.Intn(len(world.Cities))]
}

// Return true if city and all its neighbors names are in the given list  of city names
func cityAndRelationsSeen(seen map[string]bool, city *City, checked map[string]bool) bool {

	if city == nil {
		return true
	}
	// Detect loops - stop if city has already been checked for.
	if len(checked) == 0 {
		checked[city.Name] = true
	} else {
		if _, cityChecked := checked[city.Name]; cityChecked {
			return true
		}
	}

	if _, citySeen := seen[city.Name]; !citySeen {
		return false
	}

	for _, neighbor := range city.ActiveNeighbors() {
		checked[neighbor.Name] = true
		if !cityAndRelationsSeen(seen, neighbor, checked) {
			return false
		}
	}
	return true
}

// Sort type to return cities sorted descending by neighbor count, ascending by name
type byNeighborCount []*City

func (cities byNeighborCount) Len() int {
	return len(cities)
}
func (cities byNeighborCount) Swap(i, j int) {
	cities[i], cities[j] = cities[j], cities[i]
}
func (cities byNeighborCount) Less(i, j int) bool {
	log.Printf("Checking if %v < %v", cities[i], cities[j])
	if cities[i] == nil || cities[j] == nil {
		return false
	}
	iNeighbors := len(cities[i].ActiveNeighbors())
	jNeighbors := len(cities[j].ActiveNeighbors())
	if iNeighbors == jNeighbors {
		return cities[i].Name < cities[j].Name
	} else {
		return iNeighbors > jNeighbors
	}
}

// Serialize a World to string
func (world World) Serialize() string {
	seen := map[string]bool{}
	cityStrings := []string{}

	sortedByNeighborCount := []*City{}
	for _, city := range world.Cities {
		sortedByNeighborCount = append(sortedByNeighborCount, city)
	}
	sort.Sort(byNeighborCount(sortedByNeighborCount))

	for _, city := range sortedByNeighborCount {

		// If city and all its neighbors have already been seen, there's no need to serialize this city
		if cityAndRelationsSeen(seen, city, map[string]bool{}) {
			continue
		}

		// Mark city and neighbors seen
		cityStrings = append(cityStrings, city.Serialize())
		seen[city.Name] = true
		for _, neighbor := range city.ActiveNeighbors() {
			seen[neighbor.Name] = true
		}
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

	// Ensure any cities introduced as neighbors are added
	for _, neighbors := range neighborMap {
		for _, cityName := range neighbors {
			if _, ok := cityMap[cityName]; !ok {
				city := CreateCity(cityName)
				cityMap[city.Name] = &city
			}
		}
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
