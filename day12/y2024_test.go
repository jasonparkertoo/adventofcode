package day12

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay12A(t *testing.T) {
	data := utils.NewData(Example, utils.Year2024)
	
	expected := 1930
	actual := CalculateTotalPrice(data)
	
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}

func TestDay12B(t *testing.T) {
	data := utils.NewData(Challenge, utils.Year2024)
	
	expected := 1451030
	actual := CalculateTotalPrice(data)
	
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}
