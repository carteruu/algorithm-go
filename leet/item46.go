package leet

func permute11(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	seen := make([]bool, n)
	var dfs func(rs []int)
	dfs = func(rs []int) {
		if len(rs) == n {
			dst := make([]int, len(rs))
			copy(dst, rs)
			res = append(res, dst)
			return
		}
		for idx, num := range nums {
			if seen[idx] == true {
				continue
			}
			seen[idx] = true
			rs = append(rs, num)
			dfs(rs)
			rs = rs[:len(rs)-1]
			seen[idx] = false
		}
	}
	dfs([]int{})
	return res
}
