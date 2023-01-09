package leetcode_2

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
