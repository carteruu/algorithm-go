package leetcode

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := make([]int, 0, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	for i := 0; i < n*2; i++ {
		num := nums[i%n]
		for l := len(stack); l > 0 && nums[stack[l-1]%n] < num; l = len(stack) {
			ans[stack[l-1]%n] = num
			stack = stack[:l-1]
		}
		stack = append(stack, i)
	}
	return ans
}
