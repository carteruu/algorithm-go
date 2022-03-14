package leetcode

import "math/bits"

func maximumRequests(n int, requests [][]int) int {
	ans := 0
	ps := make([]int, n)
	var dfs func(idx, check, cnt int)
	dfs = func(idx, change, cnt int) {
		if idx == len(requests) {
			if change == 0 && cnt > ans {
				ans = cnt
			}
			return
		}
		//不选择
		dfs(idx+1, change, cnt)
		//选择
		req := requests[idx]
		if req[0] != req[1] {
			if ps[req[0]] == 0 {
				change++
			} else if ps[req[0]] == 1 {
				change--
			}
			if ps[req[1]] == 0 {
				change++
			} else if ps[req[1]] == -1 {
				change--
			}
		}
		ps[req[0]]--
		ps[req[1]]++
		dfs(idx+1, change, cnt+1)
		//撤销选择
		ps[req[0]]++
		ps[req[1]]--
	}
	dfs(0, 0, 0)
	return ans
}
func maximumRequests1(n int, requests [][]int) int {
	ans := 0
	var mark uint32 = 0
	for ; mark < 1<<len(requests); mark++ {
		if bits.OnesCount32(mark) <= ans {
			continue
		}
		cnt := make([]int, n)
		for j := 0; j < len(requests); j++ {
			req := requests[j]
			if mark&(1<<j) == 0 || req[0] == req[1] {
				continue
			}
			cnt[req[0]]--
			cnt[req[1]]++
		}
		check := true
		for _, c := range cnt {
			if c != 0 {
				check = false
				break
			}
		}
		if check {
			ans = bits.OnesCount32(mark)
		}
	}
	return ans
}
