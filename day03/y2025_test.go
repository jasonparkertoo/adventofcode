package day03

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var y2025ExampleData = utils.NewData(utils.Example, utils.Year2025)
var y2025ChallengeData = utils.NewData(utils.Challenge, utils.Year2025)

func TestDay3A(t *testing.T) {
	var expected int64 = 357
	actual := TotalOutputJoltage(y2025ExampleData, 2)
	
	assert.Equal(t, expected, actual)
}

func TestDay3B(t *testing.T) {
	var expected int64 = 17376
	actual := TotalOutputJoltage(y2025ChallengeData, 2)
	
	assert.Equal(t, expected, actual)
}

func TestDay3C(t *testing.T) {
	
	var expected int64 = 3121910778619
	actual := TotalOutputJoltage(y2025ExampleData, 12)
	
	assert.Equal(t, expected, actual)
}

func TestDay3D(t *testing.T) {
	
	var expected int64 = 172119830406258
	actual := TotalOutputJoltage(y2025ChallengeData, 12)
	
	assert.Equal(t, expected, actual)
}