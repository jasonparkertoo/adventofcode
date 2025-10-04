package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func Test_Day7_A(t *testing.T) {
	records, err := ReadLines(Example, "day7")
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := c.NewCalibrations(records)

	expected := 3749
	actual := c.TotalCalibrationResult(cals)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func Test_Day7_B(t *testing.T) {
	records, err := ReadLines(Challenge, "day7")
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := c.NewCalibrations(records)

	expected := 3598800864292
	actual := c.TotalCalibrationResult(cals)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func Test_Day7_C(t *testing.T) {
	records, err := ReadLines(Example, "day7")
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := c.NewCalibrations(records)

	expected := 11387
	actual := c.TotalCalibrationResultWithConcat(cals)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}

func Test_Day7_D(t *testing.T) {
	records, err := ReadLines(Challenge, "day7")
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := c.NewCalibrations(records)

	expected := 340362529351427
	actual := c.TotalCalibrationResultWithConcat(cals)

	assert.Equal(t, expected, actual, "expected %d, got %d", expected, actual)
}
