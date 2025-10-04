package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestDay2PartOneA(t *testing.T) {
	records, err := ReadLines(Example, "day2")
	if err != nil {
		panic("unable to find input data")
	}
	reports := c.NewReports(records)

	expected := 2
	actual := c.NumberOfSafeReports(false, reports)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay2PartOneB(t *testing.T) {
	records, err := ReadLines(Challenge, "day2")
	if err != nil {
		panic("unable to find input data")
	}
	reports := c.NewReports(records)
	
	expected := 606
	actual := c.NumberOfSafeReports(false, reports)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay2PartTwoA(t *testing.T) {
	records, err := ReadLines(Example, "day2")
	if err != nil {
		panic("unable to find input data")
	}
	reports := c.NewReports(records)

	expected := 4
	actual := c.NumberOfSafeReports(true, reports)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func TestDay2PartTwoB(t *testing.T) {
	records, err := ReadLines(Challenge, "day2")
	if err != nil {
		panic("unable to find input data")
	}
	reports := c.NewReports(records)

	expected := 644
	actual := c.NumberOfSafeReports(true, reports)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}
