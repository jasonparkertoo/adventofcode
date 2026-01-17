package day02

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// ToReports parses the raw input lines into a two-dimensional slice of integers.
// Each line is split by spaces and each token parsed as an int.
// Returns a slice of slices of integers representing the reports.
func ToReports(lines []string) (reports [][]int) {
	for _, line := range lines {
		var report []int
		for r := range strings.SplitSeq(line, " ") {
			n, _ := strconv.Atoi(r)
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	return
}

// isIncreasing checks if a report is strictly increasing (each value is greater than the previous).
// Returns true if the report is strictly increasing, false otherwise.
func isIncreasing(r []int) bool {
	i := 1
	for i < len(r) {
		if r[i-1] > r[i] {
			return false
		}
		i++
	}
	return true
}

// isDecreasing checks if a report is strictly decreasing (each value is less than the previous).
// Returns true if the report is strictly decreasing, false otherwise.
func isDecreasing(r []int) bool {
	i := 1
	for i < len(r) {
		if r[i-1] < r[i] {
			return false
		}
		i++
	}
	return true
}

// adjacentIsValid checks if the differences between adjacent values are valid.
// Valid differences are between 1 and 3 (inclusive).
// Returns true if all adjacent differences are valid, false otherwise.
func adjacentIsValid(r []int) bool {
	i := 1
	for i < len(r) {
		diff := r[i-1] - r[i]
		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || 3 < diff {
			return false
		}
		i++
	}
	return true
}

// trend determines the direction of a report's trend.
// Returns 1 if increasing, -1 if decreasing, 0 if no trend (same values).
func trend(r []int) int {
	if r[0] == r[1] {
		return 0
	}
	if r[0] < r[1] {
		return 1
	}
	return -1
}

// isSafe determines if a report is safe according to the rules.
// A report is safe if it is either strictly increasing or strictly decreasing
// and has valid adjacent differences (between 1 and 3).
// Returns true if the report is safe, false otherwise.
func isSafe(r []int) bool {
	if len(r) < 2 {
		return true
	}

	var monotonic bool

	switch trend(r) {
	case -1:
		monotonic = isDecreasing(r)
	case 1:
		monotonic = isIncreasing(r)
	default:
		monotonic = false
	}

	if !monotonic {
		return false
	}

	if !adjacentIsValid(r) {
		return false
	}

	return true
}

// isTolerable determines if a report is tolerable.
// A report is tolerable if it is safe or if removing any single value
// results in a safe report (i.e., it has one "bad" value that can be removed).
// Returns true if the report is tolerable, false otherwise.
func isTolerable(r []int) bool {
	if isSafe(r) {
		return true
	}

	for i := range r {
		tr := make([]int, 0, len(r)-1)
		tr = append(tr, r[:i]...)
		tr = append(tr, r[i+1:]...)

		if isSafe(tr) {
			return true
		}
	}
	return false
}

// NumberOfSafeReports counts the number of reports that are safe or tolerable based on the useDampener flag.
// If useDampener is true, a report is considered safe if it is tolerable; otherwise, it must be safe.
// The function returns the count of such reports.
func NumberOfSafeReports(useDampener bool, data *utils.Data) int {
	reports := ToReports(data.Lines())

	count := 0
	for i := range reports {
		r := reports[i]

		safe := false
		if useDampener {
			safe = isTolerable(r)
		} else {
			safe = isSafe(r)
		}

		if safe {
			count++
		}
	}
	return count
}
