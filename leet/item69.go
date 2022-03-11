package leet

func mySqrt(x int) int {
	min, max := 0, x
	ans := min
	for min <= max {
		mid := (max + min) >> 1
		product := mid * mid
		if product == x {
			return mid
		} else if product < x {
			ans = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return ans
}
func mySqrt1(x int) int {
	i := 0
	for ; i*i <= x; i++ {
	}
	return i - 1
}
