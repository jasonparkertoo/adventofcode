package day07

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay7A(t *testing.T) {
	records, err := ReadLines(Year2024, Example)
	if err != nil {
		panic(MsgPanic)
	}
	cals, _ := NewCalibrations(records)

	expected := 3749
	actual := TotalCalibrationResult(cals)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay7B(t *testing.T) {
	records, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic(MsgPanic)
	}
	cals, _ := NewCalibrations(records)

	expected := 3598800864292
	actual := TotalCalibrationResult(cals)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay7C(t *testing.T) {
	records, err := ReadLines(Year2024, Example)
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := NewCalibrations(records)

	expected := 11387
	actual := TotalCalibrationResultWithConcat(cals)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay7D(t *testing.T) {
	records, err := ReadLines(Year2024, Challenge)
	if err != nil {
		panic("unable to find input data")
	}
	cals, _ := NewCalibrations(records)

	expected := 340362529351427
	actual := TotalCalibrationResultWithConcat(cals)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
