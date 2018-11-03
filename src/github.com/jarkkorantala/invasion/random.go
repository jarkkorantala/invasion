package main

import (
	"math/rand"
)

// Interface for a random int generator
type Random interface {
	Intn(_ int) int
}

type Randomizer struct{}

// Create a randomizer
func NewRandomizer() Randomizer {
	return Randomizer{}
}

// Generate a pseudorandom int between 0 and "max"
func (randomizer Randomizer) Intn(max int) int {
	return rand.Intn(max)
}
