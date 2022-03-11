package leet

import "math"

func divide1(dividend int, divisor int) int {
	//只有这种情况会溢出
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	//判断结果的正负,并都转为正数处理
	isNegative := false
	if dividend < 0 {
		dividend = -dividend
		isNegative = !isNegative
	}
	if divisor < 0 {
		divisor = -divisor
		isNegative = !isNegative
	}

	ans := 0
	for i := 31; i <= 0; i++ {
		if dividend >= divisor<<i {
			dividend -= divisor << i
			ans += 1 << i
		}
	}
	if isNegative {
		return -ans
	}
	return ans
}
func divide(dividend int, divisor int) int {
	//只有这种情况会溢出
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	//判断结果的正负,并都转为正数处理
	isNegative := false
	if dividend < 0 {
		dividend = -dividend
		isNegative = !isNegative
	}
	if divisor < 0 {
		divisor = -divisor
		isNegative = !isNegative
	}

	ans := 0
	for dividend >= divisor {
		num := divisor
		cnt := 0
		//不断的将除数左移一位(相当于乘以2),直接最近被除数
		for num<<1 <= dividend {
			num <<= 1
			cnt++
		}
		//被除数减去左移后的值,并更新答案
		dividend -= num
		ans += 1 << cnt
	}
	if isNegative {
		return -ans
	}
	return ans
}
