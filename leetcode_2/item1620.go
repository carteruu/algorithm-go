package leetcode_2

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	ps := make(map[[2]int]int, len(towers)*radius)
	for _, tower := range towers {
		for x := maxInt(0, tower[0]-radius); x <= tower[0]+radius; x++ {
			for y := maxInt(0, tower[1]-radius); y <= tower[1]+radius; y++ {
				d := (tower[0]-x)*(tower[0]-x) + (tower[1]-y)*(tower[1]-y)
				if d > radius*radius {
					continue
				}
				ps[[2]int{x, y}] += int(float64(tower[2]) / (1 + math.Sqrt(float64(d))))
			}
		}
	}
	var ans [2]int
	var max int
	for p, c := range ps {
		if c > max || (c == max && (ans[0] < p[0] || (ans[0] == p[0] && ans[1] < p[1]))) {
			max = c
			ans = p
		}
	}
	return []int{ans[0], ans[1]}
}
