package day10

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func TestDay10A(t *testing.T) {
	assert.Equal(t, 36, TotalTrailheadScore(ex2024))
}

func TestDay10B(t *testing.T) {
	assert.Equal(t, 776, TotalTrailheadScore(ch2024))
}

func TestDay10C(t *testing.T) {
	assert.Equal(t, 81, TotalTrailheadRating(ex2024))
}

func TestDay10D(t *testing.T) {
	assert.Equal(t, 1657, TotalTrailheadRating(ch2024))
}
