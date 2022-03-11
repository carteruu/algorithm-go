package leet

import "sort"

func triangleNumber(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}

func triangleNumber4(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			idx := sort.SearchInts(nums[j+1:], v+nums[j])
			ans += idx
		}
	}
	return
}
func triangleNumber3(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			//二分查找最后一个元素的数量
			//右边: < nums[j]+nums[i]
			right := -1
			low := j + 1
			high := len(nums) - 1
			for target := nums[j] + nums[i]; low <= high; {
				mid := (high + low) / 2
				if nums[mid] < target {
					right = mid
					low = mid + 1
				} else if nums[mid] >= target {
					high = mid - 1
				}
			}
			if right != -1 {
				res += right - j
			}
		}
	}
	return res
}
func triangleNumber1(nums []int) int {
	res := 0
	sort.Ints(nums)
	var helper func(p int, ns []int)
	helper = func(p int, ns []int) {
		if p >= len(nums) {
			return
		}
		for i := p + 1; i < len(nums); i++ {
			if nums[i] <= 0 {
				continue
			}
			if len(ns) == 2 {
				if nums[i] >= ns[0]+ns[1] {
					break
				} else {
					res++
					continue
				}
			} else {
				ns = append(ns, nums[i])
				helper(i, ns)
				ns = ns[:len(ns)-1]
			}
		}
	}
	helper(-1, make([]int, 0, 3))
	return res
}
