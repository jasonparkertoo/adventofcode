package day11

import (
	"bufio"
	"strings"

	"adventofcode.dev/utils"
)

func dataTransformer(lines []string) any {
	graph := make(map[string][]string)
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		device := parts[0]
		outputs := strings.Split(parts[1], " ")

		graph[device] = outputs
	}

	return graph
}

func NumberOfDifferentPaths(d *utils.Data) int {
	// Join all lines into one input string
	var inputBuilder strings.Builder
	for _, line := range d.Lines() {
		if line != "" {
			inputBuilder.WriteString(line)
			inputBuilder.WriteString("\n")
		}
	}
	input := inputBuilder.String()

	return NumberOfPaths(input)
}

func NumberOfPaths(input string) int {
	// Build adjacency list representation of the graph
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		device := parts[0]
		outputs := strings.Split(parts[1], " ")

		graph[device] = outputs
	}

	// Find all paths from "you" to "out"
	return countPaths(graph, "you", "out")
}

func countPaths(graph map[string][]string, current, target string) int {
	// If we reached the target, we found one path
	if current == target {
		return 1
	}

	// If current device has no outputs, no paths from here
	outputs, exists := graph[current]
	if !exists || len(outputs) == 0 {
		return 0
	}

	// Count paths through each output
	totalPaths := 0
	for _, nextDevice := range outputs {
		totalPaths += countPaths(graph, nextDevice, target)
	}

	return totalPaths
}
