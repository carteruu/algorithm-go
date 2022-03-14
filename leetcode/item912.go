package leetcode

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

//快排
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := nums[r]

	idx := l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
	nums[r], nums[idx] = nums[idx], nums[r]
	quickSort(nums, l, idx-1)
	quickSort(nums, idx+1, r)
}

//归并
func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	l1 := l
	l2 := mid + 1
	mergeSort(nums, l1, mid)
	mergeSort(nums, l2, r)
	temp := make([]int, 0, r-l+1)
	for l1 <= mid && l2 <= r {
		if nums[l1] <= nums[l2] {
			temp = append(temp, nums[l1])
			l1++
		} else {
			temp = append(temp, nums[l2])
			l2++
		}
	}
	for ; l1 <= mid; l1++ {
		temp = append(temp, nums[l1])
	}
	for ; l2 <= r; l2++ {
		temp = append(temp, nums[l2])
	}
	copy(nums[l:r+1], temp)
}
