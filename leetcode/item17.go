package leetcode

import "math"

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}
	digitLetter := [][]byte{
		{},
		{},
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	ans := make([]string, 0, int(math.Pow(3.0, float64(n))))
	var dfs func(idx int, path []byte)
	dfs = func(idx int, path []byte) {
		if idx == n {
			ans = append(ans, string(path))
			return
		}
		for _, c := range digitLetter[digits[idx]-'0'] {
			path = append(path, c)
			dfs(idx+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, make([]byte, 0, n))
	return ans
}
