package leetcode_1

/**
贪心
*/
func canJump(nums []int) bool {
	//能跳到的最远位置
	m := 0
	for i := 0; i <= m && m < len(nums)-1; i++ {
		m = max(m, i+nums[i])
	}
	return m >= len(nums)-1
}
