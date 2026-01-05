package day06

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// DataTransformer splits each line into non‑empty tokens using strings.Fields,
// which removes all whitespace and returns the fields directly.
func DataTransformer(lines []string) any {
	var out [][]string
	for _, line := range lines {
		out = append(out, strings.Fields(line))
	}
	return out
}

// CalculateGrandTotal computes the grand total by performing a single pass over
// the input data. It eliminates the intermediate intVals matrix and reduces
// allocations to just two slices that accumulate column sums and products.
func CalculateGrandTotal(d *utils.Data) int {
	data := d.TransformData(DataTransformer).([][]string)

	numRows := len(data)
	if numRows == 0 {
		return 0
	}
	numCols := len(data[0])

	// Accumulators for each column.
	colSums := make([]int64, numCols)
	colProds := make([]int64, numCols)
	for i := range colProds {
		colProds[i] = 1
	}

	// Single pass over all data rows except the last (operators).
	for r := 0; r < numRows-1; r++ {
		row := data[r]
		for c, valStr := range row {
			val, _ := strconv.Atoi(valStr)
			colSums[c] += int64(val)
			colProds[c] *= int64(val)
		}
	}

	total := int64(0)
	lastRow := data[numRows-1]
	for c, op := range lastRow {
		if op == "+" {
			total += colSums[c]
		} else { // multiplication
			total += colProds[c]
		}
	}

	return int(total)
}

