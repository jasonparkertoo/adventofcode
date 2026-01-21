package day10

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025A(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 7
	actual, _ := FewestButtonPresses(ex2025)
	assert.Equal(t, expected, actual)
}

func Test2025B(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 404
	actual, _ := FewestButtonPresses(ch2025)
	assert.Equal(t, expected, actual)
}

func Test2025C(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 33
	actual, _ := CalculateScore(ex2025)
	assert.Equal(t, expected, actual)
}

func Test2025D(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 16474
	actual, _ := CalculateScore(ch2025)
	assert.Equal(t, expected, actual)
}
