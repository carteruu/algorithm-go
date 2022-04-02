package leetcode_1

import "math"

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	ans := make([]string, 0, int(math.Pow(4, float64(n))))
	letters := [][]byte{
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
	var dfs func(idx int, path []byte)
	dfs = func(idx int, path []byte) {
		if idx == n {
			ans = append(ans, string(path))
			return
		}
		for _, b := range letters[digits[idx]-'0'] {
			path[idx] = b
			dfs(idx+1, path)
		}
	}
	dfs(0, make([]byte, n))
	return ans
}
