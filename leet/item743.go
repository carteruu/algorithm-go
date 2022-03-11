package leet

func networkDelayTime(times [][]int, n int, k int) int {
	//邻接表
	m := make([][][]int, n+1)
	for _, val := range times {
		m[val[0]] = append(m[val[0]], val)
	}

	nodesTime := make([]int, n+1)
	for i := 1; i < len(nodesTime); i++ {
		nodesTime[i] = -1
	}
	nodesTime[k] = 0

	var maxDelayTime func(nodeVal, time int)
	maxDelayTime = func(nodeVal, time int) {
		children := m[nodeVal]
		for _, child := range children {
			targetVal := child[1]
			t := time + child[2]
			if nodesTime[targetVal] == -1 || t < nodesTime[targetVal] {
				nodesTime[targetVal] = t
				maxDelayTime(targetVal, t)
			}
		}
	}

	maxDelayTime(k, 0)
	maxTime := 0
	for i := 1; i < len(nodesTime); i++ {
		nodeTime := nodesTime[i]
		if nodeTime == -1 {
			return -1
		}
		if nodeTime > maxTime {
			maxTime = nodeTime
		}
	}
	return maxTime
}
