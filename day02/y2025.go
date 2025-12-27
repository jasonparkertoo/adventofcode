package day02

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type ProductIds struct {
	Ids [][]string
}

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

func SumInvalidIds(data *utils.Data) []int {
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

	return []int{sum, sum2}
}
