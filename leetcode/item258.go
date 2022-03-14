package leetcode

func addDigits1(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return addDigits(sum)
}
