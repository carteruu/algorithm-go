package leetcode_2

import "math"

func minMoves(nums []int, k int) int {
	preZeroCnt := make([]int, 0, len(nums))
	cnt0 := 0
	for _, num := range nums {
		if num == 0 {
			cnt0++
		} else {
			preZeroCnt = append(preZeroCnt, cnt0)
			cnt0 = 0
		}
	}
	preZeroSum := make([]int, len(preZeroCnt)+1)
	for i := 1; i < len(preZeroSum); i++ {
		preZeroSum[i] = preZeroSum[i-1] + preZeroCnt[i-1]
	}
	var temp int
	for i := 0; i < k-1; i++ {
		temp += preZeroCnt[i] * minInt(i+1, k-i-1)
	}
	ans := math.MaxInt
	for i := k - 1; i < len(preZeroCnt); i++ {
		temp -= preZeroSum[k/2+(i-k+1)] - preZeroSum[i-k+1]
		temp += preZeroSum[i+1] - preZeroSum[k/2+(i-k+1)+k%2]
		ans = minInt(ans, temp)
	}
	return ans
}
