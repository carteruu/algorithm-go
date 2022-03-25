package leetcode_1

import "sort"

func twoSum(nums []int, target int) []int {
	//下标切片
	idx := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx[i] = i
	}
	//将下标切片按值顺序排序
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]]-nums[idx[j]] < 0
	})
	//双指针
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[idx[l]] + nums[idx[r]]
		if sum == target {
			return []int{idx[l], idx[r]}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
func twoSum_1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
