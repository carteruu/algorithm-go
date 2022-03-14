package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)

	var twoSum func(sum int, start int) [][]int
	twoSum = func(sum int, start int) [][]int {
		rs := [][]int{}
		l := start
		r := n - 1
		for l < r {
			s := nums[l] + nums[r]
			if s < sum {
				l++
			} else if s > sum {
				r--
			} else {
				rs = append(rs, []int{-sum, nums[l], nums[r]})
				lVal := nums[l]
				for l < n && nums[l] == lVal {
					l++
				}
				rVal := nums[r]
				for r > start && nums[r] == rVal {
					r--
				}
			}
		}
		return rs
	}

	res := [][]int{}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			//排除重复元素
			continue
		}
		for _, r := range twoSum(-nums[i], i+1) {
			res = append(res, r)
		}
	}

	return res
}
