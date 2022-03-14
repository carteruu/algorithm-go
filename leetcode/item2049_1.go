package leetcode

func countHighestScoreNodes_2(parents []int) int {
	n := len(parents)
	//邻接表
	children := make([][]int, n)
	for i, parent := range parents {
		if parent < 0 {
			continue
		}
		children[parent] = append(children[parent], i)
	}
	//计算左右子节点的数量
	childCnt := make([][2]int, n)
	var dfs func(node int) int
	dfs = func(node int) int {
		if node < 0 || node >= n {
			return 0
		}
		c := 1
		for i, childNode := range children[node] {
			childCnt[node][i] = dfs(childNode)
			c += childCnt[node][i]
		}
		return c
	}
	dfs(0)

	ans := 0
	maxP := 0
	for _, cc := range childCnt {
		p := 1
		if cc[0] > 0 {
			p *= cc[0]
		}
		if cc[1] > 0 {
			p *= cc[1]
		}
		parentCnt := n - 1 - cc[0] - cc[1]
		if parentCnt > 0 {
			p *= parentCnt
		}
		if p > maxP {
			maxP = p
			ans = 1
		} else if p == maxP {
			ans++
		}
	}
	return ans
}
