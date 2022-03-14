package leetcode

import (
	math "math"
	"sort"
)

func numSquarefulPerms2(nums []int) int {
	n := len(nums)
	ans := 0
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			ans++
			return
		}
		seen := map[int]bool{}
		for i := idx; i < n; i++ {
			if seen[nums[i]] {
				continue
			}
			if idx > 0 {
				if !isSquare(nums[i] + nums[idx-1]) {
					continue
				}
			}
			nums[idx], nums[i] = nums[i], nums[idx]
			backtrack(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
			seen[nums[i]] = true
		}
	}

	backtrack(0)
	return ans
}

func isSquare(n int) bool {
	l := int(math.Sqrt(float64(n)))
	return l*l == n
}

func numSquarefulPerms(nums []int) int {
	count := 0
	var path []int
	used := map[int]bool{}
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			//fmt.Println(path)
			count++
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] {
				if !used[i-1] {
					continue
				}
			}
			if len(path) > 0 && !isSquare(path[len(path)-1]+nums[i]) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			delete(used, i)
		}
	}
	sort.Ints(nums)
	backtrack()
	return count
}
