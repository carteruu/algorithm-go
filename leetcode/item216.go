package leetcode

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var backtrack func(path []int, num, sum int)
	backtrack = func(path []int, num, sum int) {
		if sum == n && len(path) == k {
			a := make([]int, k)
			copy(a, path)
			ans = append(ans, a)
			return
		}
		if sum >= n || len(path) >= k || num > 9 {
			return
		}
		//不选择
		backtrack(path, num+1, sum)
		//选择
		path = append(path, num)
		backtrack(path, num+1, sum+num)
		path = path[:len(path)-1]
	}
	backtrack([]int{}, 1, 0)
	return ans
}
