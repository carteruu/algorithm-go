package leet

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	dp := make([]int, len(boxes)+1)
	dp[1] = 2
	trip := 2
	load := boxes[0][1]
	boxCnt := 1
	for i, j := 0, 1; j < len(boxes); j++ {
		load += boxes[j][1]
		boxCnt++
		if boxes[j][0] != boxes[j-1][0] {
			trip++
		}
		for load > maxWeight || boxCnt > maxBoxes || dp[i+1] == dp[i] {
			boxCnt--
			load -= boxes[i][1]
			i++
			if boxes[i][0] != boxes[i-1][0] {
				trip--
			}
		}
		dp[j+1] = dp[i] + trip
	}
	return dp[len(boxes)]
}
