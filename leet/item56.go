package leet

import (
	"math"
	"sort"
)

func merge12(intervals [][]int) [][]int {
	min := math.MaxInt64
	max := 0
	for _, interval := range intervals {
		if interval[0] < min {
			min = interval[0]
		}
		if interval[1] > max {
			max = interval[1]
		}
	}
	ii := make([]bool, max-min+1)
	for _, interval := range intervals {
		for i := interval[0]; i < interval[1]; i++ {
			ii[i-min] = true
		}
	}
	var ans [][]int
	s, e := -1, -1
	for i, b := range ii {
		if b {
			if s == -1 {
				s, e = i, i
			} else {
				e = i
			}
		} else {
			if s != -1 {
				ans = append(ans, []int{s + min, e + min + 1})
				s, e = -1, -1
			}
		}
	}
	if s != -1 {
		ans = append(ans, []int{s + min, e + min + 1})
		s, e = -1, -1
	}
	return ans
}
func merge1(intervals [][]int) [][]int {
	var ans [][]int
	if len(intervals) == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[j][1] < intervals[i][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	for _, interval := range intervals {
		if interval[0] > temp[1] {
			//与上一个不相交
			ans = append(ans, temp)
			temp = interval
			continue
		} else {
			if interval[1] > temp[1] {
				temp[1] = interval[1]
			}
		}
	}
	ans = append(ans, temp)
	return ans
}
