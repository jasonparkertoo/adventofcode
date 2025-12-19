package day12

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var Day12ExampleData []string
var Day12ChallengeData []string

func init() {
	Day12ExampleData, _ = ReadLines(Year2024, Example)
	Day12ChallengeData, _ = ReadLines(Year2024, Challenge)
}

func TestDay12A(t *testing.T) {
	expected := 1930
	actual := CalculateTotalPrice(Day12ExampleData)
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}

func TestDay12B(t *testing.T) {
	expected := 1451030
	actual := CalculateTotalPrice(Day12ChallengeData)
	assert.Equal(t, expected, actual, "expected %d, actual %d", expected, actual)
}
