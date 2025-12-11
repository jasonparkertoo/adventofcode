package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var Day12ExampleData []string
var Day12ChallengeData []string

func init() {
	Day12ExampleData, _ = ReadLines(Day12, PartA)
	Day12ChallengeData, _ = ReadLines(Day12, PartB)
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
