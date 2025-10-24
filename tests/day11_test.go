package tests

import (
	"strconv"
	"strings"
	"testing"

	"adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

var ExampleData []int64
var ChallengeData []int64

func init() {
	exampleData, _ := ReadLines(Example, "day11")
	ExampleData = toIntArray(exampleData[0])

	challengeData, _ := ReadLines(Challenge, "day11")
	ChallengeData = toIntArray(challengeData[0])
}

func toIntArray(str string) []int64 {
	parts := strings.Split(str, " ")
	
	out := make([]int64, len(parts))
	for i := range parts {
		n, _ := strconv.ParseInt(parts[i], 10, 64)
		out[i] = int64(n)
	}
	return out
}

func Test_Day11_Scratch(t *testing.T) {
	var expected int64 = 22
	actual := challenges.NumberOfStones(6, ExampleData)
	assert.Equal(t, expected, actual)
}

func Test_Day11_A(t *testing.T) {
	var expected int64 = 55312
	actual := challenges.NumberOfStones(25, ExampleData)
	assert.Equal(t, expected, actual)
}

func Test_Day11_B(t *testing.T) {
	var expected int64 = 202019
	actual := challenges.NumberOfStones(25, ChallengeData)
	assert.Equal(t, expected, actual)
}

func Test_Day11_C(t *testing.T) {
	var expected int64 = 239321955280205
	actual := challenges.NumberOfStones(75, ChallengeData)
	assert.Equal(t, expected, actual)
}