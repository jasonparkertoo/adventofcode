package day08

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay8A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 14
	actual := CountUniqueLocations(data)

	assert.Equal(t, expected, actual)
}

func TestDay8B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 261
	actual := CountUniqueLocations(data)

	assert.Equal(t, expected, actual)
}

func TestDay8C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 34
	actual := CountUniqueLocationsHarmonics(data)

	assert.Equal(t, expected, actual)
}

func TestDay8D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 988
	actual := CountUniqueLocationsHarmonics(data)

	assert.Equal(t, expected, actual)
}
