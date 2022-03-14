package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var dfs func(idx, t int, choicePre bool, path []int)
	dfs = func(idx, t int, choicePre bool, path []int) {
		if t == 0 {
			a := make([]int, len(path))
			copy(a, path)
			ans = append(ans, a)
			return
		}
		if idx == len(candidates) || t < candidates[idx] {
			return
		}
		//不选择
		dfs(idx+1, t, false, path)
		//选择
		if choicePre || (idx > 0 && candidates[idx-1] != candidates[idx]) {
			path = append(path, candidates[idx])
			dfs(idx+1, t-candidates[idx], true, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, target, true, make([]int, 0, len(candidates)))
	return ans
}

func combinationSum2_1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}
