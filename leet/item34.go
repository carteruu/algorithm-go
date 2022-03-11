package leet

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	//左边界
	l, r := 0, len(nums)
	for l < r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			ans[0] = mid
			r = mid
		}
	}
	//右边界
	l, r = 0, len(nums)
	for l < r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			ans[1] = mid
			l = mid + 1
		}
	}
	return ans
}
