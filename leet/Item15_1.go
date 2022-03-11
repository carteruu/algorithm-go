package leet

import "sort"

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for f, num := range nums {
		if f > 0 && num == nums[f-1] {
			continue
		}
		s, t := f+1, n-1
		for s < t {
			numS := nums[s]
			numT := nums[t]
			sum := num + numS + numT
			if sum < 0 {
				for s < t && nums[s] == numS {
					s++
				}
			} else if sum > 0 {
				for s < t && nums[t] == numT {
					t--
				}
			} else {
				ans = append(ans, []int{num, nums[s], nums[t]})
				for s < t && nums[s] == numS {
					s++
				}
				for s < t && nums[t] == numT {
					t--
				}
			}
		}
	}
	return ans
}
