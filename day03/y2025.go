package day03

import (
	"fmt"
	"strconv"

	"adventofcode.dev/utils"
)

/**
 * Finds the largest joltage by removing digits to form a number with 'keep' digits.
 * The function uses a greedy algorithm to build the result.
 *
 * @param digits - The input string of digits
 * @param keep - The number of digits to keep in the result
 * @return The resulting string with 'keep' digits forming the largest possible number
 */
func findLargestJoltage(digits string, keep int) string {
	remove := len(digits) - keep

	// Use a stack to build the result
	stack := make([]byte, 0, len(digits))

	for i := 0; i < len(digits); i++ {
		// While we can remove digits and current digit is larger than top of stack
		for remove > 0 && len(stack) > 0 && stack[len(stack)-1] < digits[i] {
			// Remove the smaller digit from stack
			stack = stack[:len(stack)-1]
			remove--
		}
		// Add current digit to stack
		stack = append(stack, digits[i])
	}

	// If we still need to remove more digits (for non-increasing sequences)
	// remove from the end
	if remove > 0 {
		stack = stack[:len(stack)-remove]
	}

	// We should have exactly 'keep' digits, but ensure it
	if len(stack) > keep {
		stack = stack[:keep]
	}

	return string(stack)
}

/**
 * Calculates the total output joltage by processing each line of data.
 * For each line, it finds the largest joltage with 'keep' digits and sums them up.
 *
 * @param data - The input data containing lines to process
 * @param keep - The number of digits to keep in each joltage calculation
 * @return The total output joltage as an int64
 */
func TotalOutputJoltage(data *utils.Data, keep int) int64 {
	var total int64 = 0

	for _, line := range data.Lines() {
		if line == "" {
			continue
		}

		// Get the largest 12-digit number from this line
		largest := findLargestJoltage(line, keep)

		// Convert to int64 and add to total
		num, err := strconv.ParseInt(largest, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			continue
		}

		total += num
	}

	return total
}
