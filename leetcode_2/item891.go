package leetcode_2

import (
	"math"
	"sort"
)

func sumSubseqWidths2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	mod := int(math.Pow10(9) + 7)
	var ans int
	var pow = 1
	for i := 0; i < n; i++ {
		ans = (ans + (nums[i]-(nums[n-1-i]))*pow) % mod
		pow = (pow * 2) % mod
	}
	return (ans + mod) % mod
}
func sumSubseqWidths1(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	mod := int64(math.Pow10(9) + 7)
	var ans int64
	var pow int64 = 1
	for i := 0; i < n; i++ {
		ans = (ans + int64(nums[i])*pow - (int64(nums[n-1-i]) * pow)) % mod
		pow = (pow * 2) % mod
	}
	return int(ans)
}
func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		pow := int(math.Pow(2, float64(n-1-i))) % mod
		ans = (ans + (nums[n-1-i]-nums[i])*pow) % mod
	}
	return (ans + mod) % mod
}
