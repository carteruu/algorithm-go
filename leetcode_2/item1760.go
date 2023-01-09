package leetcode_2

import "sort"

func minimumSize(nums []int, maxOperations int) int {
	var m int
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return sort.Search(m, func(i int) bool {
		if i == 0 {
			return false
		}
		ops := 0
		for _, num := range nums {
			ops += (num - 1) / i
		}
		return ops <= maxOperations
	})
}
