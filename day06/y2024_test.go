package day06

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay6A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 41
	actual := CountDistinctPositions(data)

	assert.Equal(t, expected, actual)
}

func TestDay6B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 5551
	actual := CountDistinctPositions(data)

	assert.Equal(t, expected, actual)
}

func TestDay6C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 6
	actual, _ := CountLoopPositions(data)

	assert.Equal(t, expected, actual)
}

func TestDay6D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 1939
	actual, _ := CountLoopPositions(data)

	assert.Equal(t, expected, actual)
}
