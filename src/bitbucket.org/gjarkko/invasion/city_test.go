package main

import (
	"testing"
)

func TestCity(test *testing.T) {
	expected := CreateCity("London")
	actual := CreateCity("London")
	if expected.String() != actual.String() {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

func TestCityNeighbours(test *testing.T) {
	washington := CreateCity("Washington DC")
	baltimore := CreateCity("Baltimore")
	philadelphia := CreateCity("Philadelphia")
	washington.SetNeighbor(East, &baltimore)
	baltimore.SetNeighbor(East, &philadelphia)

	expected := &philadelphia
	actual := washington.GetNeighbor(East).GetNeighbor(East)
	if expected.String() != actual.String() {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}

	expected = &washington
	actual = philadelphia.GetNeighbor(West).GetNeighbor(West)
	if expected.String() != actual.String() {
		test.Errorf("expected: %+v, actual: %+v", expected, actual)
	}

}
