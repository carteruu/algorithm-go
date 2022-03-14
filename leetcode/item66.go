package leetcode

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append(digits, 1)
	copy(digits[1:], digits[:n])
	digits[0] = 1
	return digits
}
