package leetcode

import (
	"math"
	"strconv"
)

func nextGreaterElement_1(n int) int {
	bs := []byte(strconv.Itoa(n))
	for i := len(bs) - 2; i >= 0; i-- {
		//从右往左遍历,找到第一个比右边邻居小的下标i(从右向左递增被打断的地方)
		if bs[i] < bs[i+1] {
			//从右往左遍历,找到第一个比下标i的数字大的下标j
			for j := len(bs) - 1; j > i; j-- {
				if bs[i] < bs[j] {
					//交换i和j
					bs[j], bs[i] = bs[i], bs[j]
					//为了形成比n稍大的组合,需要将i后面的数字逆序
					//因为原来是从右向左递增,交换后依旧是递增的,需要改成递减,才满足要求
					for l, r := i+1, len(bs)-1; l < r; {
						bs[l], bs[r] = bs[r], bs[l]
						l++
						r--
					}
					num, _ := strconv.Atoi(string(bs))
					if num < math.MinInt32 || num > math.MaxInt32 {
						return -1
					}
					return num
				}
			}
		}
	}
	return -1
}
