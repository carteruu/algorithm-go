package leetcode

func twoSum(nums []int, target int) []int {
	val2Idx := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := val2Idx[target-num]; ok {
			return []int{idx, i}
		}
		val2Idx[num] = i
	}
	return []int{}
}
