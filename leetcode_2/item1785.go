package leetcode_2

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	diff := absInt(sum - goal)
	return (diff + limit - 1) / limit
}
func minElements1(nums []int, limit int, goal int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	diff := absInt(goal - sum)
	ans := diff / limit
	if diff%limit > 0 {
		ans++
	}
	return ans
}
