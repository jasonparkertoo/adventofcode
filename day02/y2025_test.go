package day02

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay2A(t *testing.T) {
	data, _ := ReadLines(Year2025, Example)
	productIds := NewProductIds(data)
	result := productIds.SumInvalidIds()

	expected := 1227775554
	actual := result[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay2B(t *testing.T) {
	data, _ := ReadLines(Year2025, Challenge)
	productIds := NewProductIds(data)
	result := productIds.SumInvalidIds()

	expected := 15873079081
	actual := result[0]
	
	assert.Equal(t, expected, actual)
}

func TestDay2C(t *testing.T) {
	data, _ := ReadLines(Year2025, Example)
	productIds := NewProductIds(data)
	result := productIds.SumInvalidIds()

	expected := 4174379265
	actual := result[1]
	
	assert.Equal(t, expected, actual)
}

func TestDay2D(t *testing.T) {
	data, _ := ReadLines(Year2025, Challenge)
	productIds := NewProductIds(data)
	result := productIds.SumInvalidIds()

	expected := 22617871034
	actual := result[1]
	
	assert.Equal(t, expected, actual)
}