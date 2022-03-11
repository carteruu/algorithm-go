package leet

import "sort"

func canCross(stones []int) bool {
	cellStoneMap := make(map[int]int, len(stones))
	dp := make([]map[int]bool, len(stones))
	for i, val := range stones {
		cellStoneMap[val] = i
		dp[i] = make(map[int]bool)
	}
	var backtrack func(idx, step int) (res bool)
	backtrack = func(idx, step int) (res bool) {
		if idx == len(stones)-1 {
			return true
		}
		if val, ok := dp[idx][step]; ok {
			return val
		}
		defer func() { dp[idx][step] = res }()
		for i := 0; i < 3; i++ {
			curStep := step - 1 + i
			if curStep <= 0 {
				continue
			}
			val, ok := cellStoneMap[stones[idx]+curStep]
			if !ok {
				continue
			}
			if backtrack(val, curStep) {
				return true
			}
		}
		return false
	}
	return backtrack(0, 0)
}

func canCross1(stones []int) bool {
	n := len(stones)
	dp := make([]map[int]bool, n-1)
	for i := range dp {
		dp[i] = map[int]bool{}
	}
	var dfs func(int, int) bool
	dfs = func(i, lastDis int) (res bool) {
		if i == n-1 {
			return true
		}
		if res, has := dp[i][lastDis]; has {
			return res
		}
		defer func() { dp[i][lastDis] = res }()
		for curDis := lastDis - 1; curDis <= lastDis+1; curDis++ {
			if curDis > 0 {
				j := sort.SearchInts(stones, curDis+stones[i])
				if j != n && stones[j] == curDis+stones[i] && dfs(j, curDis) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, 0)
}

func canCross2(stones []int) bool {
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}
