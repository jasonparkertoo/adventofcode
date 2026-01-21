package day06

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2024A(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 41, CountDistinctPositions(ex2024))
}

func Test2024B(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 5551, CountDistinctPositions(ch2024))
}

func Test2024C(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	if actual, err := CountLoopPositions(ex2024); err != nil {
		assert.Fail(t, "FAIL!!!!!")
	} else {
		assert.Equal(t, 6, actual)
	}
}

func Test2024D(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	if actual, err := CountLoopPositions(ch2024); err != nil {
		assert.Fail(t, "FAIL!!!!!")
	} else {
		assert.Equal(t, 1939, actual)
	}
}
