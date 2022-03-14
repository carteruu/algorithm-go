package leetcode

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func singleNonDuplicate3(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if mid == 0 || mid == len(nums)-1 || (nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1]) {
			return nums[mid]
		}
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func singleNonDuplicate2(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if mid == 0 || mid == len(nums)-1 || (nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1]) {
			return nums[mid]
		}
		if (mid&1 == 0 && nums[mid] == nums[mid-1]) || (mid&1 == 1 && nums[mid] == nums[mid+1]) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
func singleNonDuplicate1(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
