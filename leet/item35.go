package leet

import "sort"

func searchInsert1(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := int(uint(r+l) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
func searchInsert2(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := int(uint(r+l) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func searchInsert3(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
