package day12

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func TestDay12A(t *testing.T) {
	assert.Equal(t, 1930, CalculateTotalPrice(ex2024))
}

func TestDay12B(t *testing.T) {
	assert.Equal(t, 1451030, CalculateTotalPrice(ch2024))
}
