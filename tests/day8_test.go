package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func Test_Day8_A(t *testing.T) {
	scan, err := ReadLines(Example, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 14
	actual := city.CountUniqueLocations()
	
	assert.Equal(t, expected, actual)
}

func Test_Day8_B(t *testing.T) {
	scan, err := ReadLines(Challenge, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 276
	actual := city.CountUniqueLocations()
	
	assert.Equal(t, expected, actual)
}

func Test_Day8_C(t *testing.T) {
	scan, err := ReadLines(Example, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 34
	actual := city.CountUniqueLocationsHarmonics()
	
	assert.Equal(t, expected, actual)
}

func Test_Day8_D(t *testing.T) {
	scan, err := ReadLines(Challenge, "day8")
	if err != nil {
		panic("unable to find input data")
	}
	city := c.City{Scan: scan}
	
	expected := 991
	actual := city.CountUniqueLocationsHarmonics()
	
	assert.Equal(t, expected, actual)
}