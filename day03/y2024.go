package day03

import (
	"regexp"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type Memory struct {
	sections []string
}

// ToMemory converts a slice of strings into a Memory struct.
//
// Parameters:
//   - lines: A slice of strings to be stored in the Memory
//
// Returns:
//   - A Memory struct containing the input lines
func ToMemory(lines []string) Memory {
	sections := make([]string, 0)
	sections = append(sections, lines...)
	return Memory{
		sections: sections,
	}
}

// ReMul is a regular expression that matches mul(), don't(), and do() patterns.
//
// This regex is used to find all occurrences of these patterns in input text.
var (
	ReMul = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
)

// multiply extracts and multiplies the two numbers from a mul() instruction.
//
// Parameters:
//   - instruction: A string containing a mul() instruction
//
// Returns:
//   - The product of the two numbers in the instruction
func multiply(instruction string) int {
	pair := strings.Split(instruction, ",")
	m1, _ := strconv.Atoi(strings.ReplaceAll(pair[0], "mul(", ""))
	m2, _ := strconv.Atoi(strings.ReplaceAll(pair[1], ")", ""))
	return m1 * m2
}

// execute processes an instruction and returns its result.
//
// Parameters:
//   - instruction: A string containing the instruction to process
//
// Returns:
//   - The result of executing the instruction
func execute(instruction string) int {
	pref := instruction[:strings.Index(instruction, "(")]
	switch pref {
	case "mul":
		return multiply(instruction)
	default:
		return 0
	}
}

// CalculateUncorrupted calculates the sum of all mul() operations in the input data.
//
// Parameters:
//   - d: A pointer to utils.Data containing input lines
//
// Returns:
//   - The sum of all valid mul() operations as an integer
func CalculateUncorrupted(d *utils.Data) int {
	input := d.Lines()
	c := ToMemory(input)

	var instructions []string
	for _, section := range c.sections {
		instructions = append(instructions, ReMul.FindAllString(section, -1)...)
	}
	sum := 0
	for _, instruction := range instructions {
		sum += execute(instruction)
	}
	return sum
}

// Calculate processes the input data and computes the sum of valid mul() operations,
// respecting do() and don't() instructions. It iterates through each section in the
// memory, finds all mul(), do(), and don't() patterns, and accumulates the product
// of mul() calls unless a preceding don't() has been encountered. The do() instruction
// resets the skip flag, allowing subsequent mul() operations to be counted again.
func Calculate(d *utils.Data) int {
	input := d.Lines()
	c := ToMemory(input)

	var instructions []string
	for _, section := range c.sections {
		instructions = append(instructions, ReMul.FindAllString(section, -1)...)
	}

	sum := 0

	skip := false
	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			skip = false
			continue
		case "don't()":
			skip = true
			continue
		default:
			if skip {
				continue
			} else {
				product := execute(instruction)
				sum += product
			}
		}
	}
	return sum
}
