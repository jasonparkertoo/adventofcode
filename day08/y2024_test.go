package day08

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay8A(t *testing.T) {
	scan, err := ReadLines(Year2024, Example)
	if err != nil {
		panic(MsgPanic)
	}
	city := City{Scan: scan}

	expected := 14
	actual := city.CountUniqueLocations()

	assert.Equal(t, expected, actual)
}

func TestDay8B(t *testing.T) {
	scan, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic(MsgPanic)
	}
	city := City{Scan: scan}

	expected := 261
	actual := city.CountUniqueLocations()

	assert.Equal(t, expected, actual)
}

func TestDay8C(t *testing.T) {
	scan, err := ReadLines(Year2024, Example)
	if err != nil {
		panic("unable to find input data")
	}
	city := City{Scan: scan}

	expected := 34
	actual := city.CountUniqueLocationsHarmonics()

	assert.Equal(t, expected, actual)
}

func TestDay8D(t *testing.T) {
	scan, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic("unable to find input data")
	}
	city := City{Scan: scan}

	expected := 988
	actual := city.CountUniqueLocationsHarmonics()

	assert.Equal(t, expected, actual)
}
