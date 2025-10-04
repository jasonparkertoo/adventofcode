package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestDay4A(t *testing.T) {
	records, err := ReadLines(Example, "day4")
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := c.NewPuzzle(records)

	expected := 18
	actual := c.Count("xmas", &puzzle)
	
	assert.Equal(t, expected, actual, "expected %s, actual %", expected, actual)
}

func TestDay4B(t *testing.T) {
	records, err := ReadLines(Challenge, "day4")
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := c.NewPuzzle(records)

	expected := 2504
	actual := c.Count("xmas", &puzzle)
	
	assert.Equal(t, expected, actual, "expected %s, actual %", expected, actual)
}

func TestDay4C(t *testing.T) {
	records, err := ReadLines(Example, "day4")
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := c.NewPuzzle(records)

	expected := 9
	actual := c.CountPattern("mas", &puzzle, c.SearchXPattern)
	
	assert.Equal(t, expected, actual, "expected %s, actual %", expected, actual)
}

func TestDay4D(t *testing.T) {
	records, err := ReadLines(Challenge, "day4")
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := c.NewPuzzle(records)

	expected := 1923
	actual := c.CountPattern("mas", &puzzle, c.SearchXPattern)
	
	assert.Equal(t, expected, actual, "expected %s, actual %", expected, actual)
}
