package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

func TestDay5A(t *testing.T) {
	records, err := ReadLines(Example, "day5")
	if err != nil {
		panic("unable to find input data")
	}
	queue := c.NewPrintQueue(records)
	
	expected := 143
	actual := c.SumMiddlePageNumbers(&queue)
	
	assert.Equal(t, expected, actual, "expected %d, got %d")
}

func TestDay5B(t *testing.T) {
	records, err := ReadLines(Challenge, "day5")
	if err != nil {
		panic("unable to find input data")
	}
	queue := c.NewPrintQueue(records)
	
	expected := 5747
	actual := c.SumMiddlePageNumbers(&queue)
	
	assert.Equal(t, expected, actual, "expected %d, got %d")
}

func TestDay5C(t *testing.T) {
	records, err := ReadLines(Example, "day5")
	if err != nil {
		panic("unable to find input data")
	}
	queue := c.NewPrintQueue(records)
	
	expected := 123
	actual := c.SumIncorrectMiddlePageNumbers(&queue)
	
	assert.Equal(t, expected, actual, "expected %d, got %d")
}

func TestDay5D(t *testing.T) {
	records, err := ReadLines(Challenge, "day5")
	if err != nil {
		panic("unable to find input data")
	}
	queue := c.NewPrintQueue(records)
	
	expected := 5502
	actual := c.SumIncorrectMiddlePageNumbers(&queue)
	
	assert.Equal(t, expected, actual, "expected %d, got %d")
}
