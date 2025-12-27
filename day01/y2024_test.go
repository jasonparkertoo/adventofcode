package day01

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay1PartOneA(t *testing.T) {
	data := NewData(Example, Year2024)
	
	expected := 11
	actual := TotalDistance(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartOneB(t *testing.T) {
	data := NewData(Challenge, Year2024)
	
	expected := 2285373
	actual := TotalDistance(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartTwoA(t *testing.T) {
	data := NewData(Example, Year2024)
	
	expected := 31
	actual := SimilarityScore(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay1PartTwoB(t *testing.T) {
	data := NewData(Challenge, Year2024)
	
	expected := 21142653
	actual := SimilarityScore(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
