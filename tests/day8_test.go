package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	scan, err := ReadLines(Example, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 14
	actual := city.CountUniqueLocations()
	
	assert.Equal(t, expected, actual)
}

func TestB(t *testing.T) {
	scan, err := ReadLines(Challenge, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 276
	actual := city.CountUniqueLocations()
	
	assert.Equal(t, expected, actual)
}

func TestC(t *testing.T) {
	scan, err := ReadLines(Example, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 34
	actual := city.CountUniqueLocationsHarmonics()
	
	assert.Equal(t, expected, actual)
}

func TestD(t *testing.T) {
	scan, err := ReadLines(Challenge, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 991
	actual := city.CountUniqueLocationsHarmonics()
	
	assert.Equal(t, expected, actual)
}