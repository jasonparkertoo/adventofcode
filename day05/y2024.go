package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type Rule struct {
	Left  int
	Right int
}

type SafetyManual struct {
	PageNumbers []int
	Rules       []Rule
}

type PrintQueue struct {
	SafetyManuals []SafetyManual
}

// parseRules parses rule definitions from lines, where each line contains a left and right page separated by '|'.
// It stops parsing when an empty line is encountered.
func parseRules(lines []string) []Rule {
	var out []Rule
	for i := range lines {
		l := lines[i]

		if i == 1175 {
			fmt.Println("found")
		}

		if l == "" {
			break
		}
		parts := strings.Split(lines[i], "|")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		out = append(out, Rule{Left: left, Right: right})
	}
	return out
}

// parsePageNumbers parses page numbers from lines, where each line contains comma‑separated integers.
// It reads lines in reverse order and stops when an empty line is found.
func parsePageNumbers(lines []string) [][]int {
	var pn [][]int
	for i := len(lines) - 1; i >= 0; i-- {
		if lines[i] == "" {
			break
		}
		var out []int
		for s := range strings.SplitSeq(lines[i], ",") {
			n, _ := strconv.Atoi(s)
			out = append(out, n)
		}
		pn = append(pn, out)
	}
	return pn
}

// ToPrintQueue transforms raw input data into a PrintQueue of SafetyManuals,
// grouping page numbers with the rules that apply to them.
func ToPrintQueue(d *utils.Data) PrintQueue {
	lines := d.Lines()
	pageNumbers := parsePageNumbers(lines)
	rules := parseRules(lines)

	var sm []SafetyManual
	for _, pn := range pageNumbers {
		var mr []Rule
		for _, rule := range rules {
			if containsPages(pn, rule) {
				mr = append(mr, rule)
			}
		}
		sm = append(sm, SafetyManual{PageNumbers: pn, Rules: mr})
	}
	return PrintQueue{SafetyManuals: sm}
}

// containsPages reports whether both endpoints of rule r appear in the given pageNumbers slice.
func containsPages(pageNumbers []int, r Rule) bool {
	return slices.Contains(pageNumbers, r.Left) && slices.Contains(pageNumbers, r.Right)
}

// conformsToRule checks that, in the page sequence p, the left endpoint of rule r appears before the right endpoint.
func conformsToRule(p []int, r Rule) bool {
	left := slices.Index(p, r.Left)
	right := slices.Index(p, r.Right)
	return left < right
}

// SumMiddlePageNumbers returns the sum of the middle page number of each SafetyManual that satisfies all its rules.
func SumMiddlePageNumbers(d *utils.Data) int {
	q := ToPrintQueue(d)
	sum := 0
	for _, s := range q.SafetyManuals {
		isValid := true
		for _, r := range s.Rules {
			if !containsPages(s.PageNumbers, r) {
				continue
			}
			if !conformsToRule(s.PageNumbers, r) {
				isValid = false
				break
			}
		}
		if isValid {
			sum += s.PageNumbers[len(s.PageNumbers)/2]
		}
	}
	return sum
}

// topologicalSort performs a topological sort on the graph represented by adjacency lists and in‑degree map.
func topologicalSort(graph map[int][]int, inDegree map[int]int) []int {
	var result []int

	for {
		// Collect all nodes with in-degree 0
		var zeroDegreeNodes []int
		for node, deg := range inDegree {
			if deg == 0 {
				zeroDegreeNodes = append(zeroDegreeNodes, node)
			}
		}

		// If none left, we're done
		if len(zeroDegreeNodes) == 0 {
			break
		}

		// Remove nodes and update neighbors
		for _, node := range zeroDegreeNodes {
			delete(inDegree, node)
			for _, neighbor := range graph[node] {
				inDegree[neighbor] = inDegree[neighbor] - 1
			}
		}

		result = append(result, zeroDegreeNodes...)
	}

	return result
}

// correctPageOrder returns a topologically sorted order of page numbers according to the safety manual's rules.
func correctPageOrder(manual SafetyManual) []int {
	// Build graph and in-degree
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, page := range manual.PageNumbers {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	// Add edges based on rules
	for _, rule := range manual.Rules {
		// Only consider rules where both pages exist in the manual
		_, leftExists := graph[rule.Left]
		_, rightExists := graph[rule.Right]
		if leftExists && rightExists {
			graph[rule.Left] = append(graph[rule.Left], rule.Right)
			inDegree[rule.Right] = inDegree[rule.Right] + 1
		}
	}

	// Topological sort
	return topologicalSort(graph, inDegree)
}

// SumIncorrectMiddlePageNumbers returns the sum of the middle page number of the first SafetyManual that violates any rule,
// after reordering its pages to satisfy the rules.
func SumIncorrectMiddlePageNumbers(d *utils.Data) int {
	q := ToPrintQueue(d)

	sum := 0
	for _, s := range q.SafetyManuals {
		for _, r := range s.Rules {
			if !conformsToRule(s.PageNumbers, r) {
				pages := correctPageOrder(s)
				sum += pages[len(pages)/2]
				break
			}
		}
	}
	return sum
}
