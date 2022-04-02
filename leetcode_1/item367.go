package leetcode_1

func isPerfectSquare(num int) bool {
	l, h := 0, num
	for l <= h {
		mid := (l + h) >> 1
		p := mid * mid
		if p == num {
			return true
		} else if p < num {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}
