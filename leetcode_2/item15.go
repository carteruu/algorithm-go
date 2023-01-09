package leetcode_2

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for a := range nums {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		t := -nums[a]
		for b, c := a+1, len(nums)-1; b < c; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[b]+nums[c] > t {
				c--
			}
			if b >= c {
				break
			}
			if nums[b]+nums[c] == t {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}
func threeSum1(nums []int) [][]int {
	m := make(map[int]struct{})
	ans := make(map[[3]int]struct{})
	sort.Ints(nums)
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			c := -a - b
			if _, ok := m[c]; ok {
				ans[[3]int{a, b, c}] = struct{}{}
			}
		}
		m[a] = struct{}{}
	}
	a := make([][]int, 0, len(ans))
	for ints := range ans {
		a = append(a, []int{ints[0], ints[1], ints[2]})
	}
	return a
}
