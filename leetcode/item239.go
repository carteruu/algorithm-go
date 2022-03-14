package leetcode

func maxSlidingWindow2(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, k)
	for i := 0; i < k-1; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		ans = append(ans, queue[0])
		if queue[0] == nums[i-k+1] {
			queue = queue[1:]
		}
	}
	return ans
}
func maxSlidingWindow1(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k-1 {
			ans = append(ans, queue[0])
			if queue[0] == nums[i-k+1] {
				queue = queue[1:]
			}
		}
	}
	return ans
}
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && queue[0] == i-k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
