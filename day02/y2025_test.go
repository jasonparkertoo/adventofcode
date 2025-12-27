package day02

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay2A(t *testing.T) {
	data := utils.NewData(Example, Year2025)
	
	expected := 1227775554
	actual := SumInvalidIds(data)[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay2B(t *testing.T) {
	data := utils.NewData(Challenge, Year2025)
	
	expected := 15873079081
	actual := SumInvalidIds(data)[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay2C(t *testing.T) {
	data := utils.NewData(Example, Year2025)

	expected := 4174379265
	actual := SumInvalidIds(data)[1]
	
	assert.Equal(t, expected, actual)
}

func TestDay2D(t *testing.T) {
	data := utils.NewData(Challenge, Year2025)
	
	expected := 22617871034
	actual := SumInvalidIds(data)[1]
	
	assert.Equal(t, expected, actual)
}