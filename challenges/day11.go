package challenges

import (
	"strconv"
	"strings"
)

func isZero(n string) bool {
	return n == "0"
}

func hasEvenNumberOfDigits(n string) bool {
	return len(n)%2 == 0
}

var strArrToString = func(strs []string) string {
	out := ""
	for _, s := range strs {
		out += s
	}
	return out
}

func splitPebble(p string) (left, right int64) {
	parts := strings.Split(p, "")
	mid := len(parts) / 2
	l, _ := strconv.ParseInt(strArrToString(parts[0:mid]), 10, 64)
	r, _ := strconv.ParseInt(strArrToString(parts[mid:]), 10, 64)

	left = int64(l)
	right = int64(r)

	return
}

func Blink(numbers string) string {
	numStrs := strings.Split(numbers, " ")

	out := ""
	for i := range numStrs {
		if len(out) > 1 {
			out += " "
		}

		str := numStrs[i]
		if isZero(str) {
			out += "1"
		} else if hasEvenNumberOfDigits(str) {
			l, r := splitPebble(str)
			out += strconv.FormatInt(l, 10)
			out += " "
			out += strconv.FormatInt(r, 10)
		} else {
			n, _ := strconv.ParseInt(str, 10, 64)
			out += strconv.FormatInt(n*2024, 10)
		}
	}
	return out
}

func countStones(stones string) int {
	return len(strings.Split(stones, " "))
}

func NumberOfStones(n int, numbers string) int {
	result := numbers
	for _ = range n {
		result = Blink(result)
	}
	return countStones(result)
}
