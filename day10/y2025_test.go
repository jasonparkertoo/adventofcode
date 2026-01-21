package day10

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	expected := 7
	actual, _ := FewestButtonPresses(ex2025)
	assert.Equal(t, expected, actual)
}

func Test2025B(t *testing.T) {
	expected := 404
	actual, _ := FewestButtonPresses(ch2025)
	assert.Equal(t, expected, actual)
}

func Test2025C(t *testing.T) {
	expected := 33
	actual, _ := CalculateScore(ex2025)
	assert.Equal(t, expected, actual)
}

func Test2025D(t *testing.T) {
	expected := 16474
	actual, _ := CalculateScore(ch2025)
	assert.Equal(t, expected, actual)
}
