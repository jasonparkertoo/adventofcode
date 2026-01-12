package day08

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025Day8A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2025)

	expected := 40
	actual := ProductOfThreeLargestCircuits(data, 10)

	assert.Equal(t, expected, actual)
}

func Test2025Day8B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2025)

	expected := 330786
	actual := ProductOfThreeLargestCircuits(data, 1000)

	assert.Equal(t, expected, actual)
}

func Test2025Day8C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2025)

	expected := 25272
	actual := ProductOfLastConnectionX(data)

	assert.Equal(t, expected, actual)
}

func Test2025Day8D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2025)

	expected := 3276581616
	actual := ProductOfLastConnectionX(data)

	assert.Equal(t, expected, actual)
}
