package leetcode_1

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	sort.Ints(nums)
	var dfs func(path []int, seen []bool)
	dfs = func(path []int, seen []bool) {
		if len(path) == n {
			ans = append(ans, append(make([]int, 0, n), path...))
			return
		}
		for i, num := range nums {
			//避免重复：当前数字与前一个数字相同时，必需前一个数字被选择，才能选择当前数字
			if seen[i] || (i > 0 && num == nums[i-1] && seen[i-1]) {
				continue
			}
			seen[i] = true
			path = append(path, num)
			dfs(path, seen)
			path = path[:len(path)-1]
			seen[i] = false
		}
	}
	dfs(make([]int, 0, n), make([]bool, n))
	return ans
}
func permuteUnique_1(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	sort.Ints(nums)
	var dfs func(path []int, seen []bool)
	dfs = func(path []int, seen []bool) {
		if len(path) == n {
			ans = append(ans, append(make([]int, 0, n), path...))
			return
		}
		//该位置，上一次选择的数字
		pre := 100000
		for i, num := range nums {
			//相同的位置，不能选择一样的数字
			if seen[i] || num == pre {
				continue
			}
			pre = num
			seen[i] = true
			path = append(path, num)
			dfs(path, seen)
			path = path[:len(path)-1]
			seen[i] = false
		}
	}
	dfs(make([]int, 0, n), make([]bool, n))
	return ans
}
