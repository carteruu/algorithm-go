package leetcode_2

import (
	"math"
)

func divide(dividend int, divisor int) int {
	//唯一会溢出的情况
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	return divide_(int32(dividend), int32(divisor))
}

func divide_(dividend int32, divisor int32) int {
	//因为负数的范围比较大，被除数和除数都转成负数计算
	if dividend > 0 {
		return -divide_(-dividend, divisor)
	}
	if divisor > 0 {
		return -divide_(dividend, -divisor)
	}
	//预处理 divisor 的倍数: 1, 2, 4, 8...直到溢出
	a := make([]int32, 0, 31)
	a = append(a, divisor)
	pows := make([]int, 0, 31)
	pow := 1
	pows = append(pows, pow)
	for divisor+divisor < divisor {
		divisor += divisor
		pow += pow
		a = append(a, divisor)
		pows = append(pows, pow)
	}
	//从最小值开始处理被除数
	ans := 0
	for i := len(a) - 1; i >= 0; i-- {
		if dividend <= a[i] {
			dividend -= a[i]
			ans += pows[i]
		}
		if dividend == 0 {
			break
		}
	}
	return ans
}
