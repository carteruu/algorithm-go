package leet

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	for i := 0; i < 1<<n; i++ {
		need := true
		path := make([]int, 0, n)
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				if j > 0 && nums[j-1] == nums[j] && i&(1<<(j-1)) == 0 {
					need = false
					break
				}
				path = append(path, nums[j])
			}
		}
		if need {
			ans = append(ans, path)
		}
	}
	return ans
}
func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	var dfs func(idx, lastChoiceIdx int, path []int)
	dfs = func(idx, lastChoiceIdx int, path []int) {
		if idx == n {
			a := make([]int, len(path))
			copy(a, path)
			ans = append(ans, a)
			return
		}
		//选择,不出现重复的条件:当前为第一个元素,或当前元素与上一个元素不同,或上一个元素被选择
		if idx == 0 || nums[idx] != nums[idx-1] || (lastChoiceIdx == idx-1) {
			path = append(path, nums[idx])
			dfs(idx+1, idx, path)
			path = path[:len(path)-1]
		}
		//不选择
		dfs(idx+1, lastChoiceIdx, path)
	}
	dfs(0, -1, make([]int, 0, n))
	return ans
}
