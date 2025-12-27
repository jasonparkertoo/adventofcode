package day11

import (
	"strconv"
	"strings"
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ExampleData []int64
var ChallengeData []int64

func init() {
	exampleData, _ := ReadLines(Year2024, Example)
	ExampleData = toIntArray(exampleData[0])

	challengeData, _ := ReadLines(Year2024, Challenge)
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

func TestDay11A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	var expected int64 = 55312
	actual := NumberOfStones(25, data)

	assert.Equal(t, expected, actual)
}

func TestDay11B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	var expected int64 = 202019
	actual := NumberOfStones(25, data)

	assert.Equal(t, expected, actual)
}

func TestDay11C(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	var expected int64 = 239321955280205
	actual := NumberOfStones(75, data)

	assert.Equal(t, expected, actual)
}
