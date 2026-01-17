package day06

import (
	"strconv"
	"strings"
	"unicode"

	"adventofcode.dev/utils"
)

// CalculateGrandTotal2 processes the vertical cephalopod math worksheet.
func CalculateGrandTotal2(d *utils.Data) int {
	lines := d.Lines()
	if len(lines) == 0 {
		return 0
	}

	// 1. Transpose the input so we can work with columns more easily.
	// Find the maximum row length to ensure we don't go out of bounds.
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Create a grid of columns
	numRows := len(lines)
	cols := make([]string, maxLen)
	for x := 0; x < maxLen; x++ {
		var colBuilder strings.Builder
		for y := range numRows {
			if x < len(lines[y]) {
				colBuilder.WriteByte(lines[y][x])
			} else {
				colBuilder.WriteByte(' ') // Pad short lines with spaces
			}
		}
		cols[x] = colBuilder.String()
	}

	var totalSum int
	var currentBlockNumbers []int
	var operator byte

	// 2. Iterate through columns to group digits into numbers and find operators.
	for _, col := range cols {
		// A full whitespace column indicates the end of a math problem block.
		if strings.TrimSpace(col) == "" {
			if len(currentBlockNumbers) > 0 && operator != 0 {
				totalSum += calculateBlock(currentBlockNumbers, operator)
			}
			// Reset for the next block
			currentBlockNumbers = nil
			operator = 0
			continue
		}

		// Check if the bottom character of this column is an operator.
		lastChar := col[len(col)-1]
		if lastChar == '+' || lastChar == '*' {
			operator = lastChar
		}

		// Extract digits from the column to form a multi-digit vertical number.
		// Note: The puzzle implies digits in one column slice form one number.
		var digitBuilder strings.Builder
		for i := 0; i < len(col)-1; i++ { // Exclude operator row
			if unicode.IsDigit(rune(col[i])) {
				digitBuilder.WriteByte(col[i])
			}
		}

		if digitBuilder.Len() > 0 {
			num, _ := strconv.ParseInt(digitBuilder.String(), 10, 32)
			currentBlockNumbers = append(currentBlockNumbers, int(num))
		}
	}

	// Handle the final block if the input doesn't end in a whitespace column.
	if len(currentBlockNumbers) > 0 && operator != 0 {
		totalSum += calculateBlock(currentBlockNumbers, operator)
	}

	return totalSum
}

// calculateBlock applies the operation to all numbers in the identified group.
func calculateBlock(nums []int, op byte) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		switch op {
		case '+':
			result += nums[i]
		case '*':
			result *= nums[i]
		}
	}
	return result
}

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
