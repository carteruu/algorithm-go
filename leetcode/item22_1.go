package leetcode

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(leftCnt, rightCnt int, path []byte)
	dfs = func(leftCnt, rightCnt int, path []byte) {
		if rightCnt == n {
			ans = append(ans, string(path))
			return
		}
		if leftCnt < n {
			path = append(path, '(')
			dfs(leftCnt+1, rightCnt, path)
			path = path[:len(path)-1]
		}
		if rightCnt < leftCnt {
			path = append(path, ')')
			dfs(leftCnt, rightCnt+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, make([]byte, 0, 2*n))
	return ans
}
