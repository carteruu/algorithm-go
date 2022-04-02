package leetcode_1

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(left, right int, bs []byte)
	dfs = func(left, right int, bs []byte) {
		if right == n {
			//左右括号数量正好为n个，加入答案
			ans = append(ans, string(bs))
			return
		}
		if left < n {
			//左括号数量 < n，可以选择左括号
			bs[left+right] = '('
			dfs(left+1, right, bs)
		}
		if left > right {
			//右括号数量 < 左括号，可以选择右括号
			bs[left+right] = ')'
			dfs(left, right+1, bs)
		}
	}
	dfs(0, 0, make([]byte, n*2))
	return ans
}
