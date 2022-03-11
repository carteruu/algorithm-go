package leet

func permute(nums []int) [][]int {
	n := len(nums)
	seen := make([]bool, n)
	ans := make([][]int, 0, n*(n-1)/2)
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == n {
			a := make([]int, n)
			copy(a, path)
			ans = append(ans, a)
			return
		}
		for i := 0; i < n; i++ {
			if seen[i] {
				continue
			}
			path = append(path, nums[i])
			seen[i] = true
			dfs(path)
			seen[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(make([]int, 0, n))
	return ans
}
