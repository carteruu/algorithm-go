package leet

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	//邻接表
	adjs := make([][]int, n)
	for i := 1; i < n; i++ {
		adjs[parents[i]] = append(adjs[parents[i]], i)
	}
	maxScore := 0
	ans := 0
	//深度优先搜索，返回子树的节点数
	var dfs func(node int) int
	dfs = func(node int) int {
		//节点数，默认是当前节点 1
		cnt := 1
		//分数，默认1
		score := 1
		//递归搜索所有子树，如果子树不为空，则需要更新当前节点的分数
		for _, child := range adjs[node] {
			childCnt := dfs(child)
			if childCnt > 0 {
				score *= childCnt
				cnt += childCnt
			}
		}
		//除了根节点，节点的分数需要乘以剩下的节点数
		if node != 0 {
			score *= n - cnt
		}
		//更新答案
		if score > maxScore {
			maxScore = score
			ans = 1
		} else if score == maxScore {
			ans++
		}
		return cnt
	}
	dfs(0)
	return ans
}
