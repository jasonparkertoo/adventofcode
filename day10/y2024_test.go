package day10

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay10A(t *testing.T) {
	lines, _ := ReadLines(Year2024, Example)
	lavaTrails := NewLavaTrails(lines)

	expected := 36
	actual := TotalTrailheadScore(lavaTrails)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10B(t *testing.T) {
	lines, _ := ReadLines(Year2024, Challenge)
	lavaTrails := NewLavaTrails(lines)

	expected := 776
	actual := TotalTrailheadScore(lavaTrails)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10C(t *testing.T) {
	lines, _ := ReadLines(Year2024, Example)
	lavaTrails := NewLavaTrails(lines)

	expected := 81
	actual := TotalTrailheadRating(lavaTrails)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10D(t *testing.T) {
	lines, _ := ReadLines(Year2024, Challenge)
	lavaTrails := NewLavaTrails(lines)

	expected := 1657
	actual := TotalTrailheadRating(lavaTrails)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
