package leet

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			newTarget := target - nums[i] - nums[j]
			k := j + 1
			l := len(nums) - 1
			for k < l {
				if nums[k]+nums[l] < newTarget {
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				} else if nums[k]+nums[l] > newTarget {
					l--
					for k < l && l < len(nums)-1 && nums[l] == nums[l+1] {
						l--
					}
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				}
			}
		}
	}
	return ans
}
