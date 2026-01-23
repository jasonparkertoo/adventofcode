package day12

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025A(t *testing.T) {
	data, _ := utils.NewData(utils.Example, utils.Year2025)
	assert.Equal(t, 2, CountValidRegions(data))
}

func Test2025B(t *testing.T) {
	data, _ := utils.NewData(utils.Challenge, utils.Year2025)
	assert.Equal(t, 427, CountValidRegions(data))
}