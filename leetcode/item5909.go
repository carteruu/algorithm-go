package leetcode

func minimumTime(n int, relations [][]int, time []int) int {
	n = len(time)
	inDeg := make([]map[int]struct{}, n)
	outDeg := make([][]int, n)
	for i := 0; i < n; i++ {
		inDeg[i] = make(map[int]struct{})
	}
	for _, relation := range relations {
		inDeg[relation[1]-1][relation[0]-1] = struct{}{}
		outDeg[relation[0]-1] = append(outDeg[relation[0]-1], relation[1]-1)
	}
	//idx
	inDeg0 := make([]int, 0, n)
	for i, ins := range inDeg {
		if len(ins) == 0 {
			inDeg0 = append(inDeg0, i)
		}
	}
	ans := 0
	//课程的实际时间
	aTime := make([]int, n)
	copy(aTime, time)
	for len(inDeg0) > 0 {
		course := inDeg0[0]
		inDeg0 = inDeg0[1:]
		for _, next := range outDeg[course] {
			//修改下一个课的实际时间
			if aTime[course]+time[next] > aTime[next] {
				aTime[next] = aTime[course] + time[next]
			}
			delete(inDeg[next], course)
			if len(inDeg[next]) == 0 {
				inDeg0 = append(inDeg0, next)
			}
		}
		if aTime[course] > ans {
			ans = aTime[course]
		}
	}
	return ans
}
