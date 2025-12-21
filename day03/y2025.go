package day03

import (
	"fmt"
	"strconv"

	"adventofcode.dev/utils"
)

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

