package leetcode_1

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k)
	queue := make([]int, 0, k)
	for i, num := range nums {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			if queue[0] < i-k {
				queue = queue[1:]
			}
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
