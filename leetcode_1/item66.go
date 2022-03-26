package leetcode_1

func plusOne(digits []int) []int {
	s := 1
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		d := digits[i] + s
		digits[i] = d % 10
		s = d / 10
		if s == 0 {
			//没有进位直接返回
			return digits
		}
	}
	ans := make([]int, n+1, n+1)
	ans[0] = 1
	return ans
}
func plusOne_2(digits []int) []int {
	ans := make([]int, len(digits)+1)
	s := 1
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + s
		ans[i] = d % 10
		s = d / 10
	}
	if s > 0 {
		ans[0] = s
	} else {
		ans = ans[1:]
	}
	return ans
}
func plusOne_1(digits []int) []int {
	s := 1
	for i := len(digits) - 1; i >= 0 && s > 0; i-- {
		d := digits[i] + s
		digits[i] = d % 10
		s = d / 10
	}
	if s > 0 {
		ans := []int{1}
		ans = append(ans, digits...)
		return ans
	} else {
		return digits
	}
}
