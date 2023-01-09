package leetcode_2

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	if truckSize == 0 {
		return 0
	}
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	ans := 0
	for _, boxType := range boxTypes {
		if truckSize <= boxType[0] {
			ans += truckSize * boxType[1]
			break
		}
		ans += boxType[0] * boxType[1]
		truckSize -= boxType[0]
	}
	return ans
}
