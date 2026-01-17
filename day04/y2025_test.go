package day04

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	assert.Equal(t, 13, CountAccessible(ex2025))
}

func Test2025B(t *testing.T) {
	assert.Equal(t, 1416, CountAccessible(ch2025))
}

func Test2025C(t *testing.T) {
	assert.Equal(t, 43, CountRemovable(ex2025))
}

func Test2025D(t *testing.T) {
	assert.Equal(t, 9086, CountRemovable(ch2025))
}
