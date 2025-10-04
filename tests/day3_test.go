package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestPartOneA(t *testing.T) {
	records, err := ReadLines(Example, "day3")
	if err != nil {
		panic("unable to find input data")
	}
	instructions := c.NewInstructions(records)

	expected := 161
	actual := c.CalculateUncorrupted(instructions)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestPartOneB(t *testing.T) {
	records, err := ReadLines(Challenge, "day3")
	if err != nil {
		panic("unable to find input data")
	}
	instructions := c.NewInstructions(records)

	expected := 156388521
	actual := c.CalculateUncorrupted(instructions)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestPartTwoA(t *testing.T) {
	records, err := ReadLines(Example, "day3_2")
	if err != nil {
		panic("unable to find input data")
	}
	instructions := c.NewInstructions(records)

	expected := 48
	actual := c.Calculate(instructions)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestPartTwoB(t *testing.T) {
	records, err := ReadLines(Challenge, "day3")
	if err != nil {
		panic("unable to find input data")
	}
	instructions := c.NewInstructions(records)

	expected := 75920122
	actual := c.Calculate(instructions)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}
