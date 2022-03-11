package leet

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var rec func(path []int, visited []bool)
	rec = func(path []int, visited []bool) {
		if len(path) == len(nums) {
			ans = append(ans, path)
			return
		}
		for i, num := range nums {
			if visited[i] {
				//已使用过的元素,跳过
				continue
			}
			if i > 0 && visited[i-1] && num == nums[i-1] {
				//已选择了上一个元素,且当前元素和上一个元素值相同时,跳过,避免重复排列
				continue
			}
			a := make([]int, len(path))
			copy(a, path)
			a = append(a, num)
			visited[i] = true
			rec(a, visited)
			visited[i] = false
		}
	}
	rec(make([]int, 0, len(nums)), make([]bool, len(nums)))
	return ans
}

func permuteUnique1(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
