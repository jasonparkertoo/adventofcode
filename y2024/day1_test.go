package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay1PartOneA(t *testing.T) {
	records, err := ReadLines(Day1, PartA)
	if err != nil {
		panic(MsgPanic)
	}
	locationIds := NewLocationIds(records)

	expected := 11
	actual := TotalDistance(locationIds)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartOneB(t *testing.T) {
	records, err := ReadLines(Day1, PartB)
	if err != nil {
		panic(MsgPanic)
	}
	locationIds := NewLocationIds(records)

	expected := 2285373
	actual := TotalDistance(locationIds)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartTwoA(t *testing.T) {
	records, err := ReadLines(Day1, PartA)
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := NewLocationIds(records)

	expected := 31
	actual := SimilarityScore(locationIds)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartTwoB(t *testing.T) {
	records, err := ReadLines(Day1, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	locationIds := NewLocationIds(records)

	expected := 21142653
	actual := SimilarityScore(locationIds)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
