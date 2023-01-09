package leetcode_2

import "strconv"

func isPalindrome(x int) bool {
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

func isPalindrome1(x int) bool {
	//下面的代码不能处理个位数为0的数，所以这里先排除掉
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	p := 0
	for x > p {
		p = p*10 + x%10
		x /= 10
	}
	return x == p || x == p/10
}

func isPalindrome2(x int) bool {
	a, b := 0, x
	for b > 0 {
		a = a*10 + b%10
		b /= 10
	}
	return a == x
}
