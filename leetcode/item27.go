package leetcode

func removeElement(nums []int, val int) int {
	slow := 0
	for _, num := range nums {
		if num != val {
			nums[slow] = num
			slow++
		}
	}
	return slow
}
