package offer

import "math"

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	res := math.MaxInt64
	for left < right {
		mid := (right-left)/2 + left
		if numbers[right] > numbers[mid] {
			right = mid
		} else if numbers[right] < numbers[mid] {
			left = mid + 1
		} else {
			if numbers[left] > numbers[mid] {
				right = mid
			} else if numbers[left] < numbers[mid] {
				res = numbers[left]
				left = mid + 1
			} else {
				left++
				right--
			}
		}
	}
	if numbers[left] < res {
		res = numbers[left]
	}
	return res
}
