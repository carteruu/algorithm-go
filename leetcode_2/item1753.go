package leetcode_2

func maximumScore(a, b, c int) int {
	sum := a + b + c
	maxVal := maxInt(maxInt(a, b), c)
	if sum < maxVal*2 {
		return sum - maxVal
	} else {
		return sum / 2
	}
}

func maximumScore1(a int, b int, c int) int {
	var ans int
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	ans = minInt(c-b, a)
	a -= ans
	c -= ans

	ans += a
	b -= a / 2
	c -= a/2 + (a & 1)

	ans += minInt(b, c)
	return ans
}
