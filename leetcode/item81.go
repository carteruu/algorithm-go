package leetcode

func search(nums []int, target int) bool {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target || nums[l] == target || nums[r] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] < target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
func search2(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return true
		}
		if nums[l] < nums[mid] {
			//前半部分有序
			if target >= nums[l] && target < nums[mid] {
				return binarySearch(nums[l:mid], target)
			} else {
				l = mid + 1
			}
		} else if nums[r] > nums[mid] {
			//后半部分有序
			if target > nums[mid] && target <= nums[r] {
				return binarySearch(nums[mid+1:r+1], target)
			} else {
				r = mid - 1
			}
		} else {
			return search(nums[l:mid], target) || search(nums[mid+1:r+1], target)
		}
	}
	return false
}
func binarySearch(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func search1(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
