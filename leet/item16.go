package leet

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans int
	diff := math.MaxInt64
	n := len(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum == target {
				return target
			} else if sum > target {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
			d := abs(sum - target)
			if d < diff {
				diff = d
				ans = sum
			}
		}
	}
	return ans
}
