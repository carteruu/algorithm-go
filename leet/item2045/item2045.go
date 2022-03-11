package item2045

import "math"

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adjs := make([][]int, n+1)
	//邻接表
	for _, edge := range edges {
		adjs[edge[0]] = append(adjs[edge[0]], edge[1])
		adjs[edge[1]] = append(adjs[edge[1]], edge[0])
	}
	//到达某个节点的最短和第二短时间
	nodesTime := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		nodesTime[i] = [2]int{math.MaxInt64, math.MaxInt64}
	}
	//广度优先遍历,q队列[节点编号,所花时间]
	q := make([][2]int, 1, n)
	q[0] = [2]int{1, 0}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		//计算到达下个节点的所需时间
		t := node[1]
		//计算当前红绿灯,当前为红灯时,需要等待变为绿灯
		if t/change&1 == 1 {
			t += change - t%change
		}
		//加上穿过边的时间
		t += time
		//遍历下一个可达的节点
		for _, nextNode := range adjs[node[0]] {
			if t < nodesTime[nextNode][1] {
				//当t<第二短时间
				if t < nodesTime[nextNode][0] {
					//且t<最短时间,修改最短时间和第二短时间,并将该节点和所需时间加入队列
					nodesTime[nextNode][1], nodesTime[nextNode][0] = nodesTime[nextNode][0], t
					q = append(q, [2]int{nextNode, t})
				} else if t > nodesTime[nextNode][0] {
					//且t>最短时间,修改第二短时间,并将该节点和所需时间加入队列
					//第二短必需严格大于最短时间,所以这里忽略等于
					nodesTime[nextNode][1] = t
					q = append(q, [2]int{nextNode, t})
				}
			}
		}
	}
	return nodesTime[n][1]
}
