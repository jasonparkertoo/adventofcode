package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

var Day12ExampleData []string
var Day12ChallengeData []string

func init() {
	Day12ExampleData, _ = ReadLines(Example, "day12")
	Day12ChallengeData, _ = ReadLines(Challenge, "day12")
}

func Test_Day12_A(t *testing.T) {
	expected := 1930
	actual := c.CalculateTotalPrice(Day12ExampleData)
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}

func Test_Day12_B(t *testing.T) {
	expected := 1930
	actual := c.CalculateTotalPrice(Day12ChallengeData)
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}
