package day03

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var y2025ExampleData = utils.NewData(utils.Example, utils.Year2025)
var y2025ChallengeData = utils.NewData(utils.Challenge, utils.Year2025)

func TestDay3A(t *testing.T) {
	results := TotalOutputVoltage(y2025ExampleData)
	
	expected := 357
	actual := results[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay3B(t *testing.T) {
	results := TotalOutputVoltage(y2025ChallengeData)
	
	expected := 17376
	actual := results[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay3C(t *testing.T) {
	results := TotalOutputVoltage(y2025ExampleData)
	
	expected := 3121910778619
	actual := results[1]
	
	assert.Equal(t, expected, actual)
}

func TestDay3D(t *testing.T) {
	results := TotalOutputVoltage(y2025ChallengeData)
	
	expected := 0
	actual := results[1]
	
	assert.Equal(t, expected, actual)
}