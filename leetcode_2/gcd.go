package leetcode_2

func gcd(a, b int) int {
	for c := a % b; c != 0; c = a % b {
		a, b = b, c
	}
	return b
}
func gcd1(a, b int) int {
	c := a % b
	if c == 0 {
		return b
	}
	return gcd(b, c)
}
