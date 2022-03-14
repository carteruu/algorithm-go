package leetcode

func combinationSum1(candidates []int, target int) [][]int {
	var ans [][]int
	var dfs func(t, idx int, path []int)
	dfs = func(t, idx int, path []int) {
		if t == 0 {
			a := make([]int, len(path))
			copy(a, path)
			ans = append(ans, a)
			return
		}
		if idx == len(candidates) {
			return
		}
		if t >= candidates[idx] {
			path = append(path, candidates[idx])
			dfs(t-candidates[idx], idx, path)
			path = path[:len(path)-1]
		}

		dfs(t, idx+1, path)
	}
	dfs(target, 0, []int{})
	return ans
}
