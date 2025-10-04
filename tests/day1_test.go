package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestDay1PartOneA(t *testing.T) {
	records, err := ReadLines(Example, "day1")
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := c.NewLocationIds(records)

	expected := 11
	actual := c.TotalDistance(locationIds)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay1PartOneB(t *testing.T) {
	records, err := ReadLines(Challenge, "day1")
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := c.NewLocationIds(records)

	expected := 2285373
	actual := c.TotalDistance(locationIds)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay1PartTwoA(t *testing.T) {
	records, err := ReadLines(Example, "day1")
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := c.NewLocationIds(records)

	expected := 31
	actual := c.SimilarityScore(locationIds)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay1PartTwoB(t *testing.T) {
	records, err := ReadLines(Challenge, "day1")
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := c.NewLocationIds(records)

	expected := 21142653
	actual := c.SimilarityScore(locationIds)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}
