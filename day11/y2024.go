package day11

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

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

func powerOfTen(n int) int64 {
	half := n / 2
	var p int64 = 1
	for range half {
		p *= 10
	}
	return p
}

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

func NumberOfStones(depth int64, d *utils.Data) int64 {
	nums := d.Formatted(DataFormatter).([]int64)
	
	memo := make(map[[2]int64]int64)
	var total int64
	for _, n := range nums {
		total += blink(n, depth, memo)
	}
	return total
}
