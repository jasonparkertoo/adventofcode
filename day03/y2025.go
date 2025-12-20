package day03

import (
	"strconv"

	"adventofcode.dev/utils"
)

func TotalOutputVoltage(data *utils.Data) []int {
	sum := 0
	for _, bank := range data.Lines() {
		left, right := 0, 0
		for i := 0; i < len(bank); i++ {
			num, _ := strconv.Atoi(bank[i : i+1])
			if num > left {
				if i == len(bank)-1 {
					right = num
				} else {
					left = num
					right = 0
				}
			} else {
				if num > right {
					right = num
				}
			}
		}
		sum += (left * 10) + right
	}
	return []int{sum}
}

