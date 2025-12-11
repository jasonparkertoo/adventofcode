package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay2PartOneA(t *testing.T) {
	records, err := ReadLines(Day2, PartA)
	if err != nil {
		panic(MsgPanic)
	}
	reports := NewReports(records)

	expected := 2
	actual := NumberOfSafeReports(false, reports)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartOneB(t *testing.T) {
	records, err := ReadLines(Day2, PartB)
	if err != nil {
		panic(MsgPanic)
	}
	reports := NewReports(records)

	expected := 606
	actual := NumberOfSafeReports(false, reports)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartTwoA(t *testing.T) {
	records, err := ReadLines(Day2, PartA)
	if err != nil {
		panic("unable to find input data")
	}
	reports := NewReports(records)

	expected := 4
	actual := NumberOfSafeReports(true, reports)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartTwoB(t *testing.T) {
	records, err := ReadLines(Day2, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	reports := NewReports(records)

	expected := 644
	actual := NumberOfSafeReports(true, reports)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
