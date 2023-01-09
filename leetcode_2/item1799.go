package leetcode_2

import "math/bits"

func maxScore(nums []int) int {
	m := len(nums)
	dp := make([]int, 1<<m)
	gcds := make([][]int, m)
	for i := 0; i < len(nums); i++ {
		gcds[i] = make([]int, m)
		for j := i + 1; j < len(nums); j++ {
			gcds[i][j] = gcd(nums[i], nums[j])
		}
	}

	all := 1 << m
	for s := 1; s < all; s++ {
		t := bits.OnesCount(uint(s))
		//需要是偶数
		if (t & 1) != 0 {
			continue
		}
		for i := 0; i < m; i++ {
			if ((s >> i) & 1) != 0 {
				for j := i + 1; j < m; j++ {
					if ((s >> j) & 1) != 0 {
						dp[s] = maxInt(dp[s], dp[s^(1<<i)^(1<<j)]+t/2*gcds[i][j])
					}
				}
			}
		}
	}
	return dp[all-1]
}

//超时
func maxScore1(nums []int) int {
	gcds := make(map[int]int, len(nums)*len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			gcds[i*10+j] = gcd(nums[i], nums[j])
		}
	}

	var dfs func(op int, seen map[int]struct{}) int
	dfs = func(op int, seen map[int]struct{}) int {
		var m int
		for i := 0; i < len(nums); i++ {
			if _, ok := seen[i]; ok {
				continue
			}
			seen[i] = struct{}{}
			for j := 0; j < len(nums); j++ {
				if _, ok := seen[j]; ok {
					continue
				}
				seen[j] = struct{}{}
				m = maxInt(m, op*gcds[i*10+j]+dfs(op+1, seen))
				delete(seen, j)
			}
			delete(seen, i)
		}
		return m
	}
	return dfs(1, make(map[int]struct{}, len(nums)))
}
