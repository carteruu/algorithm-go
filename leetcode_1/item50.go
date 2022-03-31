package leetcode_1

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	ans := myPow(x, n/2)
	ans *= ans
	if n&1 == 1 {
		ans *= x
	}
	if neg {
		return 1 / ans
	} else {
		return ans
	}
}
func myPow_1(x float64, n int) float64 {
	var dfs func(p int) float64
	dfs = func(p int) float64 {
		if p == 0 {
			return 1
		}
		ans := myPow(x, p/2)
		ans *= ans
		if p&1 == 1 {
			//奇数
			ans *= x
		}
		return ans
	}
	if n >= 0 {
		return dfs(n)
	} else {
		return 1 / dfs(-n)
	}
}
