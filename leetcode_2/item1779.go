package leetcode_2

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	minD := math.MaxInt
	ans := -1
	for i, point := range points {
		if point[0] != x && point[1] != y {
			continue
		}
		d := absInt(point[0] - x + point[1] - y)
		if d < minD {
			d = minD
			ans = i
		}
	}
	return ans
}
