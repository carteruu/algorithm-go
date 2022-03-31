package leetcode_1

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, n)
	var dfs func(idx int, l int)
	dfs = func(idx int, l int) {
		if idx == n {
			item := make([]int, l)
			copy(item, path[:l])
			ans = append(ans, item)
			return
		}
		dfs(idx+1, l)
		path[l] = nums[idx]
		dfs(idx+1, l+1)
	}
	dfs(0, 0)
	return ans
}
func subsets_1(nums []int) [][]int {
	var ans [][]int
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		if idx == len(nums) {
			item := make([]int, len(path))
			copy(item, path)
			ans = append(ans, item)
			return
		}
		dfs(idx+1, path)
		dfs(idx+1, append(path, nums[idx]))
	}
	dfs(0, nil)
	return ans
}
