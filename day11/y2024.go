package day11

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// DataFormatter transforms a slice of raw string lines into a slice of int64 values.
// Each line may contain multiple integers separated by spaces.
func DataFormatter(data []string) any {
	out := []int64{}
	for _, str := range data {
		parts := strings.SplitSeq(str, " ")
		for p := range parts {
			n, _ := strconv.ParseInt(p, 10, 64)
			out = append(out, n)
		}
	}
	return out
}

// numberOfDigits returns the number of decimal digits in n.
// It correctly handles zero and negative values.
func numberOfDigits(n int64) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	d := 0
	for n > 0 {
		n /= 10
		d++
	}
	return d
}

// powerOfTen returns 10 raised to the power of (n / 2).
// It is used to split a number into its high and low halves.
func powerOfTen(n int) int64 {
	half := n / 2
	var p int64 = 1
	for range half {
		p *= 10
	}
	return p
}

// processNumber splits n into its left and right halves when it has an even
// number of digits.  For odd‑digit numbers it returns n multiplied by 2024
// as the left part and -1 to indicate that there is no right part.
func processNumber(n int64) (left, right int64) {
	if n == 0 {
		return 1, -1
	}
	d := numberOfDigits(n)
	if d%2 == 0 {
		p := powerOfTen(d)
		return n / p, n % p
	}
	return n * 2024, -1
}

// blink recursively evaluates a value based on the structure defined by
// processNumber, up to a given depth, and memoizes intermediate results
// to avoid redundant computations.  It returns 1 for depth 0 and combines
// results of left and right parts otherwise.
func blink(n, depth int64, memo map[[2]int64]int64) int64 {
	key := [2]int64{n, int64(depth)}
	if v, ok := memo[key]; ok {
		return v
	}
	if depth == 0 {
		return 1
	}

	l, r := processNumber(n)
	var out int64
	if r == -1 {
		out = blink(l, depth-1, memo)
	} else {
		out = blink(l, depth-1, memo) + blink(r, depth-1, memo)
	}

	memo[key] = out
	return out
}

// NumberOfStones computes the total result of applying blink to each number
// in the provided data set up to the specified depth.
func NumberOfStones(depth int64, d *utils.Data) int64 {
	nums := d.TransformData(DataFormatter).([]int64)

	memo := make(map[[2]int64]int64)
	var total int64
	for _, n := range nums {
		total += blink(n, depth, memo)
	}
	return total
}
