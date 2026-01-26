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

// NumberOfDifferentPathsWithBoth counts all distinct paths from "svr" to "out"
// that visit both "dac" and "fft" (in any order). It uses a memoized depth‑first
// traversal of the graph built from the input data.
func NumberOfDifferentPathsWithBoth(d *utils.Data) int {
	// Build graph
	graph := make(map[string][]string)
	for _, line := range d.Lines() {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		device := parts[0]
		outputs := strings.Split(parts[1], " ")
		graph[device] = outputs
	}

	// Memoization map: key is [current, seenDAC, seenFFT] as array
	memo := make(map[MemoKey]int)
	initialKey := MemoKey{"svr", false, false}
	return countPathsWithBothMemo(graph, initialKey, memo)
}

type MemoKey struct {
	Current string
	SeenDAC bool
	SeenFFT bool
}

// countPathsWithBothMemo recursively counts paths from current to out, tracking whether
// the path has encountered dac and fft, with memoization.
func countPathsWithBothMemo(graph map[string][]string, key MemoKey, memo map[MemoKey]int) int {
	// Check if result is already computed
	if result, exists := memo[key]; exists {
		return result
	}

	// Update flags if current node is one of the required nodes.
	if key.Current == "dac" {
		key.SeenDAC = true
	}
	if key.Current == "fft" {
		key.SeenFFT = true
	}

	// If we've reached the target, check flags.
	if key.Current == "out" {
		if key.SeenDAC && key.SeenFFT {
			memo[key] = 1
			return 1
		}
		memo[key] = 0
		return 0
	}

	outputs, exists := graph[key.Current]
	if !exists || len(outputs) == 0 {
		memo[key] = 0
		return 0
	}

	total := 0
	for _, next := range outputs {
		nextKey := MemoKey{next, key.SeenDAC, key.SeenFFT}
		total += countPathsWithBothMemo(graph, nextKey, memo)
	}

	memo[key] = total
	return total
}
