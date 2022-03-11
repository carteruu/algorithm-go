package leet

import (
	"math"
	"sort"
)

func minimumDifference2(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-k+1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}
func minimumDifference1(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	ans := math.MaxInt
	for right := 0; right < len(nums); right++ {
		if right-left+1 == k {
			diff := nums[right] - nums[left]
			if diff < ans {
				ans = diff
			}
			left++
		}
	}
	return ans
}
