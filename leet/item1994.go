package leet

var MOD = int(1e9) + 7
var ps = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var es = []int{4, 9, 16, 25}

func numberOfGoodSubsets(nums []int) int {
	cnt := [31]int{}
	for _, num := range nums {
		//排除不满足条件的数（包含平方因子）
		need := true
		for _, e := range es {
			if num%e == 0 {
				need = false
				break
			}
		}
		if !need {
			continue
		}
		cnt[num]++
	}
	dp := make([]int, 1<<len(ps))
	//初始化空集为1
	dp[0] = 1
	for num := 2; num <= 30; num++ {
		if cnt[num] == 0 {
			continue
		}
		mark := 0
		for j, p := range ps {
			if num%p == 0 {
				mark |= 1 << j
			}
		}
		for state := 1<<len(ps) - 1; state >= 0; state-- {
			//如果当前状态包含当前数字的所有质因子，则当前状态 = 由前一个状态状态（去掉当前数字的质因子） * 当前数字的数量
			//前一个状态为 state ^ mark
			if state&mark == mark {
				prevState := dp[state^mark]
				dp[state] = (dp[state] + prevState*cnt[num]) % MOD
			}

			// 如果当前数字的质因子和当前状态完全不同，则可以更新下一个状态的值 = 原值+当前状态数量*当前数字数量
			//下一个状态为当前状态 | 当前数字的质因子
			//if state&mark == 0 {
			//	newState := dp[state|mark]
			//	newState = (newState + dp[state]*cnt[num]) % MOD
			//}
		}
	}
	ans := 0
	//累加，不包含空集
	for i := 1; i < len(dp); i++ {
		ans = (ans + dp[i]) % MOD
	}
	//数字1的影响，n个1，则有n²种组合
	for i := 0; i < cnt[1]; i++ {
		ans = ans * 2 % MOD
	}
	return ans
}
func numberOfGoodSubsets1(nums []int) int {
	cnt := make([]int, 31)
	ans := 0
	for _, num := range nums {
		need := true
		for _, e := range es {
			if num%e == 0 {
				need = false
				break
			}
		}
		if need {
			cnt[num]++
		}
	}
	var dfs func(num int, multiple int, nn int)
	dfs = func(num int, multiple int, nn int) {
		if num > 30 {
			ans = (ans + nn) % MOD
			return
		}
		//不选择
		dfs(num+1, multiple, nn)
		//选择
		if cnt[num] > 0 && gcd(multiple, num) == 1 {
			dfs(num+1, multiple*num, nn*cnt[num]%MOD)
		}
	}
	for i := 2; i <= 30; i++ {
		dfs(i, i, cnt[i])
	}
	for i := 0; i < cnt[1]; i++ {
		ans = ans * 2 % MOD
	}
	return ans
}
