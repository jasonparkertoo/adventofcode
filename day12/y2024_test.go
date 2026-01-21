package day12

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay12A(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 1930, CalculateTotalPrice(ex2024))
}

func TestDay12B(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 1451030, CalculateTotalPrice(ch2024))
}
