package day02

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// ProductIds holds the parsed ranges from the input. Each entry in
// Ids is a two‑element slice containing the string representations
// of the start and end of a range.
type ProductIds struct {
	Ids [][]string
}

// ToProductIds parses the raw input lines into a ProductIds value.
// It expects the first line of data to contain comma‑separated
// ranges such as "1-3,5-7". Each range is split on the dash into
// a two‑element string slice. The function returns a pointer to the
// resulting ProductIds struct.
func ToProductIds(data []string) *ProductIds {
	ranges := data[0]
	parts := strings.Split(ranges, ",")

	ids := [][]string{}
	for _, p := range parts {
		ids = append(ids, strings.Split(p, "-"))
	}
	return &ProductIds{
		Ids: ids,
	}
}

// CheckInvalid determines whether the provided numeric string
// satisfies either of the two invalidity criteria described in the
// package comment. The first returned value, isInv, is true if the
// string's first half equals the second half. The second returned
// value, isInv2, is true if the string can be decomposed into a
// repeated pattern of equal length.
func CheckInvalid(num string) (bool, bool) {
	isInv := false
	if len(num)%2 == 0 {
		halfLen := len(num) / 2
		if num[:halfLen] == num[halfLen:] {
			isInv = true
		}
	}

	isInv2 := false
	n := len(num)
	for k := 2; k <= n; k++ {
		if n%k != 0 {
			continue
		}
		patternLen := n / k
		pattern := num[:patternLen]
		valid := true
		for i := 1; i < k; i++ {
			start := i * patternLen
			if num[start:start+patternLen] != pattern {
				valid = false
				break
			}
		}
		if valid {
			isInv2 = true
			break
		}
	}

	return isInv, isInv2
}

// SumInvalidIds iterates over all product identifiers, expands the
// ranges specified in the input data, and sums identifiers that
// are invalid according to the two rules. The first sum includes
// IDs that satisfy the first rule (palindromic halves), and the
// second sum includes IDs that satisfy the second rule
// (repeating patterns). The function returns a slice containing
// the two sums: []int{sum, sum2}.
func SumInvalidIds(data *utils.Data) (int, int) {
	lines := data.Lines()
	p := ToProductIds(lines)

	sum := 0
	sum2 := 0
	for _, id := range p.Ids {
		l, _ := strconv.Atoi(id[0])
		r, _ := strconv.Atoi(id[1])

		for i := l; i <= r; i++ {
			n := strconv.Itoa(i)
			isInv, isInv2 := CheckInvalid(n)
			if isInv {
				sum += i
			}
			if isInv2 {
				sum2 += i
			}
		}
	}

	return sum, sum2
}
