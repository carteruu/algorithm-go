package leetcode_2

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := make(map[int]int)
	for _, num := range arr {
		cnt[num]++
	}
	times := make([]int, 0, len(cnt))
	for _, time := range cnt {
		times = append(times, time)
	}
	sort.Ints(times)
	for i, time := range times {
		if k >= time {
			k -= time
		} else {
			return len(times) - i
		}
	}
	return 0
}
