package day06

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 41, CountDistinctPositions(ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 5551, CountDistinctPositions(ch2024))
}

func Test2024C(t *testing.T) {
	if actual, err := CountLoopPositions(ex2024); err != nil {
		assert.Fail(t, "FAIL!!!!!")
	} else {
		assert.Equal(t, 6, actual)
	}
}

func Test2024D(t *testing.T) {
	if actual, err := CountLoopPositions(ch2024); err != nil {
		assert.Fail(t, "FAIL!!!!!")
	} else {
		assert.Equal(t, 1939, actual)
	}
}
