package offer

func subsets1(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		if idx == n {
			a := make([]int, len(path))
			copy(a, path)
			ans = append(ans, a)
			return
		}
		dfs(idx+1, path)
		path = append(path, nums[idx])
		dfs(idx+1, path)
		path = path[:len(path)-1]
	}
	dfs(0, make([]int, 0, n))
	return ans
}

func subsets(nums []int) [][]int {
	n := len(nums)
	//有1<<n种可能
	ans := make([][]int, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		set := make([]int, 0, n)
		for i, v := range nums {
			//当每一位为1是,代表该元素需要加入
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, set)
	}
	return ans
}
