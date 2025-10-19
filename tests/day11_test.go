package tests

import (
	"testing"

	"adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

var ExampleData string
var ChallengeData string

func init() {
	exampleData, _ := ReadLines(Example, "day11")
	ExampleData = exampleData[0]

	challengeData, _ := ReadLines(Challenge, "day11")
	ChallengeData = challengeData[0]
}

func Test_Day11_A(t *testing.T) {
	expected := 55312
	actual := challenges.NumberOfStones(25, ExampleData)
	assert.Equal(t, expected, actual)
}

func Test_Day11_B(t *testing.T) {
	expected := 202019
	actual := challenges.NumberOfStones(25, ChallengeData)
	assert.Equal(t, expected, actual)
}