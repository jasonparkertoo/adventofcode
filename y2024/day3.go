package y2024

import (
	"regexp"
	"strconv"
	"strings"
)

type Memory struct {
	sections []string
}

func NewInstructions(lines []string) Memory {
	sections := make([]string, 0)
	sections = append(sections, lines...)
	return Memory{
		sections: sections,
	}
}

var (
	ReMul = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don\'t\(\)|do\(\)`)
)

func multiply(instruction string) int {
	pair := strings.Split(instruction, ",")
	m1, _ := strconv.Atoi(strings.ReplaceAll(pair[0], "mul(", ""))
	m2, _ := strconv.Atoi(strings.ReplaceAll(pair[1], ")", ""))
	return m1 * m2
}

func execute(instruction string) int {
	pref := instruction[:strings.Index(instruction, "(")]
	switch pref {
	case "mul":
		return multiply(instruction)
	default:
		return 0
	}
}

func CalculateUncorrupted(c Memory) int {
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

func Calculate(c Memory) int {
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
