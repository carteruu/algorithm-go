package leetcode_2

func orderOfLargestPlusSign(n int, mines [][]int) int {
	zeroSet := make(map[int]struct{})
	for _, mine := range mines {
		zeroSet[(mine[0]<<10)|mine[1]] = struct{}{}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < minInt(minInt(i, n-1-i), minInt(j, n-1-j)); k++ {
				if _, ok := zeroSet[(i<<10)|j]; ok {
					continue
				}
			}
		}
	}
	return 0
}
