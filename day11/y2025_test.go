package day11

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	
	expected := 5
	actual := NumberOfDifferentPaths(ex2024)
	
	assert.Equal(t, expected, actual)
}

func TestB(t *testing.T) {
	ex2024, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	
	expected := 764
	actual := NumberOfDifferentPaths(ex2024)
	
	assert.Equal(t, expected, actual)
}