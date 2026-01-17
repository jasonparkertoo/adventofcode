package day08

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	assert.Equal(t, 40, ProductOfThreeLargestCircuits(ex2025, 10))
}

func Test2025B(t *testing.T) {
	assert.Equal(t, 330786, ProductOfThreeLargestCircuits(ch2025, 1000))
}

func Test2025C(t *testing.T) {
	assert.Equal(t, 25272, ProductOfLastConnectionX(ex2025))
}

func Test2025D(t *testing.T) {
	assert.Equal(t, 3276581616, ProductOfLastConnectionX(ch2025))
}
