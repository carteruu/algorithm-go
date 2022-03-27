package leetcode_1

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	var quickSort func(l, r, qty int)
	quickSort = func(l, r, qty int) {
		if l >= r {
			return
		}
		p := arr[r]
		pos := l
		for i := l; i < r; i++ {
			if arr[i] < p {
				arr[i], arr[pos] = arr[pos], arr[i]
				pos++
			}
		}
		arr[pos], arr[r] = arr[r], arr[pos]
		//前半段的数字个数
		num := pos - l + 1
		if num == qty {
			//如果已经将数组分为两部分，且前半段为k个数字，结束递归
			return
		} else if num > qty {
			//如果前半段的数字个数 > qty，只需要继续排序前半段
			quickSort(l, pos-1, qty)
		} else {
			//如果前半段的数字个数 < qty，则需要排序后半段，且 qty = qty - 前半段数字个数
			quickSort(pos+1, r, qty-num)
		}
	}
	quickSort(0, len(arr)-1, k)
	return arr[:k]
}
func getLeastNumbers_1(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
