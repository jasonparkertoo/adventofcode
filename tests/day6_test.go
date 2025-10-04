package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestDay6A(t *testing.T) {
	records, err := ReadLines(Example, "day6")
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := c.Generate(records)
	path, _ := c.Explore(m)
	
	expected := 41
	actual := len(path)
	
	assert.Equal(t, expected, actual)
}

func TestDay6B(t *testing.T) {
	records, err := ReadLines(Challenge, "day6")
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := c.Generate(records)
	path, _ := c.Explore(m)
	
	expected := 5551
	actual := len(path)
	
	assert.Equal(t, expected, actual)
}

func TestDay6C(t *testing.T) {
	records, err := ReadLines(Example, "day6")
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := c.Generate(records)
	
	expected := 6
	actual, _ := c.CountLoopPositions(m)
	
	assert.Equal(t, expected, actual)
}

func TestDay6D(t *testing.T) {
	records, err := ReadLines(Challenge, "day6")
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := c.Generate(records)
	
	expected := 1939
	actual, _ := c.CountLoopPositions(m)
	
	assert.Equal(t, expected, actual)
}
