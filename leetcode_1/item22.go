package leetcode_1

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(left, right int, bs []byte)
	dfs = func(left, right int, bs []byte) {
		if right == n {
			ans = append(ans, string(bs))
			return
		}
		if left < n {
			bs[left+right] = '('
			dfs(left+1, right, bs)
		}
		if left > right {
			bs[left+right] = ')'
			dfs(left, right+1, bs)
		}
	}
	dfs(0, 0, make([]byte, n*2))
	return ans
}
