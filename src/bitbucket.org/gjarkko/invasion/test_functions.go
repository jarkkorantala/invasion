package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Fail "test" if string "actual" doesn't equal to "expected"
func assertStringEqual(test *testing.T, expected string, actual string) {
	if expected != actual {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

// Fail "test" if "actual" doesn't stringify to "expected"
func assertStringerEqual(test *testing.T, expected fmt.Stringer, actual fmt.Stringer) {
	assertStringEqual(test, expected.String(), actual.String())
}

func assertPanics(test *testing.T) {
	if recovered := recover(); recovered == nil {
		test.Errorf("The code did not panic")
	}
}

// Fail "test" if "actual" isn't a nil pointer to a City
func assertNilCityReference(test *testing.T, actual *City) {
	if actual != nil {
		test.Errorf("expected: nil, actual: %+v", actual)
	}
}

type MockRandomizer struct{}

// Create a predictable randomizer for testing
func NewMockRandomizer() MockRandomizer {
	return MockRandomizer{}
}

// Return 1 instead of a random integer
func (randomizer MockRandomizer) Intn(max int) int {
	return 1
}

// Init random by pseudorandom seed (current time)
func initRand() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

// Init random to a given seed (useful for testing)
func initRandToSeed(seed int) {
	rand.Seed(int64(seed))
}
