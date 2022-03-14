package offer

func majorityElement(nums []int) int {
	n := 0
	cnt := 0
	for _, num := range nums {
		if n == num {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				n = num
				cnt = 1
			}
		}
	}
	return n
}
