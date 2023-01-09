package leetcode_2

func removeElement(nums []int, val int) int {
	ans := 0
	for _, num := range nums {
		if num != val {
			nums[ans] = num
			ans++
		}
	}
	return ans
}
