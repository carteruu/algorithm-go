package leetcode

func findDuplicate11(nums []int) int {
	slow, fast := 0, 0
	n := len(nums)
	for fast < n && nums[fast] < n {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fast = 0
			for slow != fast {
				slow = nums[slow]
				fast = nums[fast]
			}
			return slow
		}
	}
	return -1
}
func findDuplicate2(nums []int) int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				return num
			}
		}
	}
	return -1
}
func findDuplicate12(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i]-1 {
			swapIdx := nums[i] - 1
			if nums[swapIdx] == nums[i] {
				return nums[i]
			}
			nums[i], nums[swapIdx] = nums[swapIdx], nums[i]
		}
	}
	return -1
}

func findDuplicate3(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func findDuplicate4(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}
