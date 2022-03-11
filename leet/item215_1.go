package leet

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	idx := n - k

	var partition func(l, r int) int
	var quickSelect func(l, r int) int
	partition = func(l, r int) int {
		ri := rand.Int()%(r-l+1) + l
		nums[ri], nums[r] = nums[r], nums[ri]
		pivot := nums[r]
		i := l
		for j := l; j < r; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]
		return i
	}
	quickSelect = func(l, r int) int {
		q := partition(l, r)
		if q == idx {
			return nums[q]
		} else if q < idx {
			return quickSelect(q+1, r)
		} else {
			return quickSelect(l, q-1)
		}
	}

	return quickSelect(0, n-1)
}
