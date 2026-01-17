package day07

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// Equation represents a single calibration equation consisting of a target result
// and a slice of numbers that can be combined with + or * to achieve that result.
type Equation struct {
	result  int
	numbers []int
}

// Calibrations is a collection of calibration equations.
type Calibrations struct {
	equations []Equation
}

// ToCalibrations converts a slice of input lines into a Calibrations
// structure.  Each line is expected to contain a sequence of numbers
// followed by the target result, separated by spaces.  Empty lines are
// ignored.  The function returns an error if any parsing operation fails.
func ToCalibrations(lines []string) (*Calibrations, error) {
	var equations []Equation
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")

		res := 0
		nums := make([]int, 0)

		for _, s := range parts {
			if strings.Contains(s, ":") {
				n, _ := strconv.Atoi(strings.ReplaceAll(s, ":", ""))
				res = n
			} else {
				n, _ := strconv.Atoi(s)
				nums = append(nums, n)
			}
		}

		equations = append(equations, Equation{res, nums})
	}
	return &Calibrations{equations: equations}, nil
}

// isValid recursively determines whether the sequence of numbers can
// produce the target result by inserting either + or * between adjacent
// numbers.  The recursion explores both addition and multiplication
// at each step.  The index parameter indicates the current position
// in the numbers slice, and currentValue holds the value computed so
// far from the left side of the expression.
func isValid(numbers []int, target int, index int, currentValue int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	next := numbers[index]

	// Addition
	if isValid(numbers, target, index+1, currentValue+next) {
		return true
	}

	// Multiplication
	if isValid(numbers, target, index+1, currentValue*next) {
		return true
	}

	return false
}

// TotalCalibrationResult calculates the sum of the target results for
// all equations that can be satisfied by inserting + or * between the
// numbers.  It ignores equations that cannot reach the target.
func TotalCalibrationResult(d *utils.Data) int {
	lines := d.Lines()
	c, _ := ToCalibrations(lines)
	sum := 0
	for _, eq := range c.equations {
		if isValid(eq.numbers, eq.result, 1, eq.numbers[0]) {
			sum += eq.result
		}
	}
	return sum
}

// isValidWithConcat is similar to isValid but additionally allows
// concatenating the current value with the next number (e.g. 1 and 2
// can become 12).  This function explores all three possibilities
// (addition, multiplication, and concatenation) at each step.
func isValidWithConcat(numbers []int, target int, index int, currentValue int) bool {
	if index == len(numbers) {
		return currentValue == target
	}

	nextNumber := numbers[index]

	if isValidWithConcat(numbers, target, index+1, currentValue+nextNumber) {
		return true
	}

	if isValidWithConcat(numbers, target, index+1, currentValue*nextNumber) {
		return true
	}

	concatVal, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, nextNumber))
	return isValidWithConcat(numbers, target, index+1, concatVal)
}

// TotalCalibrationResultWithConcat calculates the sum of the target
// results for all equations that can be satisfied when +, *, or
// concatenation is allowed between numbers.  Equations that cannot
// reach the target are excluded from the sum.
func TotalCalibrationResultWithConcat(data *utils.Data) int {
	c, _ := ToCalibrations(data.Lines())

	sum := 0
	for _, eq := range c.equations {
		// index and currentValue are ignored in this approach
		if isValidWithConcat(eq.numbers, eq.result, 1, eq.numbers[0]) {
			sum += eq.result
		}
	}
	return sum
}
