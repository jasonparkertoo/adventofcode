package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

type Calibrations struct {
	equations []Equation
}

func NewCalibrations(lines []string) (*Calibrations, error) {
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

func TotalCalibrationResult(c *Calibrations) int {
	sum := 0
	for _, eq := range c.equations {
		if isValid(eq.numbers, eq.result, 1, eq.numbers[0]) {
			sum += eq.result
		}
	}
	return sum
}

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

func TotalCalibrationResultWithConcat(c *Calibrations) int {
	sum := 0
	for _, eq := range c.equations {
		// index and currentValue are ignored in this approach
		if isValidWithConcat(eq.numbers, eq.result, 1, eq.numbers[0]) {
			sum += eq.result
		}
	}
	return sum
}
