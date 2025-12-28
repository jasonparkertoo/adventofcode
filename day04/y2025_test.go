package day04

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay4A(t *testing.T) {
	d := NewData(Example, Year2025)

	expected := 13
	actual := CountAccessible(d)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4B(t *testing.T) {
	d := NewData(Challenge, Year2025)

	expected := 1416
	actual := CountAccessible(d)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}