package leetcode_2

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	ans := 1.0
	c := x
	for n > 0 {
		if n&1 == 1 {
			ans *= c
		}
		c *= c
		n >>= 1
	}
	return ans
}
func myPow1(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	pow := myPow(x, n/2)
	if n&1 == 1 {
		return pow * pow * x
	}
	return pow * pow
}
