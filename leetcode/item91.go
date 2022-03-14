package leetcode

func numDecodings11(s string) int {
	n := len(s)
	//初始
	dp1 := 1
	dp2 := 1
	//状态转移
	for i := 1; i <= n; i++ {
		dp3 := 0
		if s[i-1] != '0' && (i == n || s[i] != '0') {
			dp3 += dp2
		}
		if i > 1 {
			if (s[i-2] == '1') || (s[i-2] == '2' && (s[i-1] <= '6')) {
				dp3 += dp1
			}
		}
		dp1 = dp2
		dp2 = dp3
	}
	return dp2
}
func numDecodings21(s string) int {
	n := len(s)
	//dp数组表示,字符长度为dp下标值时,解码方法的个数
	dp := make([]int, n+1)
	//初始
	dp[0] = 1
	//状态转移
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' && (i == n || s[i] != '0') {
			dp[i] += dp[i-1]
		}
		if i > 1 {
			if (s[i-2] == '1') || (s[i-2] == '2' && (s[i-1] <= '6')) {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n]
}
func numDecodings12(s string) int {
	n := len(s)
	//备忘录
	m := make(map[int]int)
	//递归
	var backtrack func(idx int) int
	backtrack = func(idx int) int {
		if idx == n {
			return 1
		}
		if val, ok := m[idx]; ok {
			return val
		}
		ans := 0
		if s[idx] != '0' && (idx == n-1 || s[idx+1] != '0') {
			//选择一位
			ans += backtrack(idx + 1)
		}
		if idx < n-1 {
			if (s[idx] == '1') || (s[idx] == '2' && (s[idx+1] <= '6')) {
				//选择两位
				ans += backtrack(idx + 2)
			}
		}
		m[idx] = ans
		return ans
	}
	return backtrack(0)
}
