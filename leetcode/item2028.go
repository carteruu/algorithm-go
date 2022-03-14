package leetcode

func missingRolls(rolls []int, mean int, n int) []int {
	//m次投掷的和
	sumM := 0
	for _, roll := range rolls {
		sumM += roll
	}
	//缺失的n次透支的和
	sumN := mean*(len(rolls)+n) - sumM
	if sumN < n || sumN > 6*n {
		return []int{}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	sumN -= n
	i := 0
	for sumN > 0 {
		if sumN > 5 {
			ans[i] += 5
			sumN -= 5
			i++
		} else {
			ans[i] += sumN
			break
		}
	}
	return ans
}
func missingRolls2(rolls []int, mean int, n int) []int {
	//m次投掷的和
	sumM := 0
	for _, roll := range rolls {
		sumM += roll
	}
	//缺失的n次透支的和
	sumN := mean*(len(rolls)+n) - sumM
	if sumN < 0 {
		return []int{}
	}
	//dp存剩余值的可选范围(闭区间),下标零代表上一次剩余的值,即已确定的sumN
	dp := make([][2]int, n+1)
	dp[0] = [2]int{sumN, sumN}
	//从前往后遍历,从上一次已确定的值,计算第1,2,3...次投掷后剩余值范围
	for i := 1; i < n; i++ {
		//剩余值下限,即当次投出6
		//剩余值上限,即当次投出1
		dp[i] = [2]int{dp[i-1][0] - 6, dp[i-1][1] - 1}
		if dp[i][1] < 0 {
			//最大的剩余值小于0,不存在答案
			return []int{}
		}
	}
	//最后一次投掷的剩余值(0)
	dp[n] = [2]int{0, 0}
	//从后往前遍历,计算倒数第1,2,3...次投掷后剩余值范围
	//再和从前往后计算的求交集
	for i := n - 1; i > 0; i-- {
		//剩余值下限,即当次投出1
		//剩余值上限,即当次投出6
		low := dp[i+1][0] + 1
		high := dp[i+1][1] + 6
		//取交集
		if low > dp[i][0] {
			dp[i][0] = low
		}
		if high < dp[i][1] {
			dp[i][1] = high
		}
		//区间为空,不存在答案
		if dp[i][0] > dp[i][1] {
			return []int{}
		}
	}
	//随便取一个可行值
	ans := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		//上一次投掷后的剩余值-当次投掷后的剩余值=投掷的点数
		ans = append(ans, dp[i-1][0]-dp[i][0])
	}
	return ans
}
func missingRolls1(rolls []int, mean int, n int) []int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	//目标值
	sum = mean*(len(rolls)+n) - sum
	ans := make([]int, 0, n)

	var dfs func(idx, mm int) bool
	dfs = func(idx, mm int) bool {
		nn := sum / (n - idx)
		if mm < 0 || nn < 1 || nn > 6 {
			return false
		}
		if idx == n {
			if mm == 0 {
				return true
			}
			return false
		}
		for i := 1; i <= 6; i++ {
			ans = append(ans, i)
			if dfs(idx+1, mm-i) {
				return true
			}
			ans = ans[:len(ans)-1]
		}
		return false
	}
	dfs(0, sum)
	return ans
}
