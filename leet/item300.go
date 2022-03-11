package leet

import (
	"algorithm/pkg"
)

func lengthOfLIS11(nums []int) int {
	len := len(nums)
	arr := make([][]int, len)

	res := 0
	for i := 0; i < len; i++ {
		arr[i] = make([]int, len)
		arr[i][i] = 1
		next := nums[i]
		for j := i + 1; j < len; j++ {
			if nums[j] > next {
				arr[i][j] = arr[i][j-1] + 1
				next = nums[j]
			} else {
				arr[i][j] = arr[i][j-1]
			}
		}
		res = pkg.Max(res, arr[i][len-1])
	}
	return res
}
