package day11

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func TestDay11A(t *testing.T) {
	assert.Equal(t, int64(55312), NumberOfStones(25, ex2024))
}

func TestDay11B(t *testing.T) {
	assert.Equal(t, int64(202019), NumberOfStones(25, ch2024))
}

func TestDay11C(t *testing.T) {
	assert.Equal(t, int64(239321955280205), NumberOfStones(75, ch2024))
}
