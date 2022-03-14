package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	revertedNumber := 0
	for temp > 0 {
		revertedNumber = 10*revertedNumber + temp%10
		temp /= 10
	}
	return x == revertedNumber
}
func isPalindrome2(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
