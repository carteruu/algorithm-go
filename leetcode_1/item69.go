package leetcode_1

import "math"

/**
牛顿迭代法快速寻找平方根
*/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	for c, x0 := float64(x), float64(x); ; {
		xi := 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			return int(xi)
		}
		x0 = xi
	}
}

/**
二分，O(log n)
*/
func mySqrt_2(x int) int {
	l, h, ans := 0, x, 0
	for l <= h {
		mid := (l + h) >> 1
		p := mid * mid
		if p == x {
			return mid
		} else if p < x {
			ans = mid
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return ans
}

/**
O(n)
*/
func mySqrt_1(x int) int {
	ans := 0
	for ans*ans <= x {
		ans++
	}
	return ans - 1
}
