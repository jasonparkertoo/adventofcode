package day05

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay5A(t *testing.T) {
	records, err := ReadLines(Year2024, Example)
	if err != nil {
		panic(MsgPanic)
	}
	queue := NewPrintQueue(records)

	expected := 143
	actual := SumMiddlePageNumbers(&queue)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5B(t *testing.T) {
	records, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic(MsgPanic)
	}
	queue := NewPrintQueue(records)

	expected := 5747
	actual := SumMiddlePageNumbers(&queue)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5C(t *testing.T) {
	records, err := ReadLines(Year2024, Example)
	if err != nil {
		panic("unable to find input data")
	}
	queue := NewPrintQueue(records)

	expected := 123
	actual := SumIncorrectMiddlePageNumbers(&queue)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5D(t *testing.T) {
	records, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic("unable to find input data")
	}
	queue := NewPrintQueue(records)

	expected := 5502
	actual := SumIncorrectMiddlePageNumbers(&queue)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
