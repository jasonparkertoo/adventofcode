package day04

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 18, Count("xmas", ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 2504, Count("xmas", ch2024))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 9, CountPattern("mas", ex2024, SearchXPattern))
}

func Test2024D(t *testing.T) {
	assert.Equal(t, 1923, CountPattern("mas", ch2024, SearchXPattern))
}
