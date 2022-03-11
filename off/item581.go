package off

import (
	"math"
	"sort"
)

func findUnsortedSubarray3(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	left, right := 0, len(nums)-1
	for nums[left] == numsSorted[left] {
		left++
	}
	for nums[right] == numsSorted[right] {
		right--
	}
	return right - left + 1
}
func findUnsortedSubarray4(nums []int) int {
	n := len(nums)
	l, r := -1, -1
	minn, maxx := math.MaxInt64, math.MinInt64
	for i := 0; i < n; i++ {
		if maxx > nums[i] {
			//从左到右,没有递增,记录最大右边界
			r = i
		} else {
			maxx = nums[i]
		}

		if minn < nums[n-1-i] {
			//从右到左,没有递减,记录最小左边界
			l = n - 1 - i
		} else {
			minn = nums[n-1-i]
		}
	}
	if l == r {
		//左边界等于右边界,说明原数组有序
		return 0
	}
	return r - l + 1
}
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	stackL := make([]int, 0, n)
	stackR := make([]int, 0, n)
	l, r := n, -1
	for i := 0; i < n; i++ {
		for size := len(stackL); size > 0 && nums[stackL[size-1]] > nums[i]; size = len(stackL) {
			if stackL[size-1] < l {
				l = stackL[size-1]
			}
			stackL = stackL[:size-1]
		}
		stackL = append(stackL, i)

		for size := len(stackR); size > 0 && nums[stackR[size-1]] < nums[n-1-i]; size = len(stackR) {
			if stackR[size-1] > r {
				r = stackR[size-1]
			}
			stackR = stackR[:size-1]
		}
		stackR = append(stackR, n-1-i)
	}
	if l >= r {
		return 0
	}
	return r - l + 1
}
func findUnsortedSubarray1(nums []int) int {
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	sort.Ints(sortNums)
	l, r := 0, len(nums)-1
	for ; l < len(nums) && nums[l] == sortNums[l]; l++ {
	}
	for ; r >= l && nums[r] == sortNums[r]; r-- {
	}
	return r - l + 1
}
