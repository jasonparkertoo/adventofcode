package day05

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025A(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	assert.Equal(t, 3, CountFreshIds(ex2025))
}

func Test2025B(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	assert.Equal(t, 896, CountFreshIds(ch2025))
}

func Test2025C(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	assert.Equal(t, int64(14), NumberOfFreshRangeIds(ex2025))
}

func Test2025D(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	assert.Equal(t, int64(346240317247002), NumberOfFreshRangeIds(ch2025))
}
