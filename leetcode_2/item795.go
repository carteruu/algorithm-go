package leetcode_2

func numSubarrayBoundedMax1(nums []int, left int, right int) int {
	count := func(lower int) (res int) {
		cur := 0
		for _, x := range nums {
			if x <= lower {
				cur++
			} else {
				cur = 0
			}
			res += cur
		}
		return
	}
	return count(right) - count(left-1)
}
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	last1, last2, res := -1, -1, 0
	for i, x := range nums {
		if left <= x && x <= right {
			last1 = i
		} else if x > right {
			last2 = i
			last1 = -1
		}
		if last1 != -1 {
			res += last1 - last2
		}
	}
	return res
}
