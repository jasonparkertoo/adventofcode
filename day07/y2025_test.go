package day07

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025Day7A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2025)
	
	expected := 21
	actual := CountBeamSplits(data)
	
	assert.Equal(t, expected, actual)
}

func Test2025Day7B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2025)
	
	expected := 1533
	actual := CountBeamSplits(data)
	
	assert.Equal(t, expected, actual)
}