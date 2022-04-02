package leetcode_1

func majorityElement(nums []int) int {
	val, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			val = num
			cnt = 1
			continue
		}
		if val == num {
			cnt++
		} else {
			cnt--
		}
	}
	return val
}
