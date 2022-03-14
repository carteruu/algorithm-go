package leetcode

import (
	"math"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	strs := strings.Split(s, " ")
	pre := math.MinInt64
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err == nil {
			if num <= pre {
				return false
			}
			pre = num
		}
	}
	return true
}
