package challenges

import (
	"strconv"
	"strings"
)

type Data struct {
	reports [][]int
}

func NewReports(lines []string) Data {
	var reports [][]int

	for _, line := range lines {
		var report []int
		for r := range strings.SplitSeq(line, " ") {
			n, _ := strconv.Atoi(r)
			report = append(report, n)
		}
		reports = append(reports, report)
	}

	return Data{
		reports: reports,
	}
}

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

func adjacentIsValid(r []int) bool {
	i := 1
	for i < len(r) {
		diff := r[i-1]-r[i]
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

func trend(r []int) int {
	if r[0] == r[1] {
		return 0
	}
	if (r[0] < r[1]) {
		return 1
	}
	return -1
}

func isSafe(r []int) bool {
	if len(r) < 2 {
		return true
	}
	
	var monotonic bool
	
	switch (trend(r)) {
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

func isSafe2(r []int) bool {
	left, right := 0, 1
	
	increasing := r[left] < r[right]
	
	numBad := 0
	for right < len(r) {
		diff := r[left] - r[right]

		if increasing && diff < 0 {
			numBad++
		}
		
		if !increasing && diff > 0 {
			numBad++
		}
		
		if numBad > 0 {
			return false
		}
	}
	
	return true
}

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

func NumberOfSafeReports(useDampener bool, c Data) int {
	count := 0
	for i := range c.reports {
		r := c.reports[i]
		
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
