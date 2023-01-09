package leetcode_2

import "math"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	row := []float64{float64(poured)}
	for i := 1; i < query_row; i++ {
		nextRow := make([]float64, i+1)
		for j, v := range row {
			if v > 1 {
				nextRow[j] += (v - 1) / 2
				nextRow[j+1] += (v - 1) / 2
			}
		}
		row = nextRow
	}
	return math.Min(1, row[query_glass])
}
