package lcp

import "math"

func minimalExecTime1(root *TreeNode) float64 {
	//求每个节点及其子节点,串行的执行时间和最大可并行的执行时间
	var dfs func(node *TreeNode) (float64, float64)
	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0.0, 0.0
		}
		left1, left2 := dfs(node.Left)
		right1, right2 := dfs(node.Right)

		tc := float64(node.Val) + left1 + right1
		if left1 < right1 {
			left1, right1 = right1, left1
			left2, right2 = right2, left2
		}

		var pc float64
		if left1-2*left2 <= right1 {
			pc = (left1 + right1) / 2
		} else {
			pc = right1 + left2
		}
		return tc, pc
	}
	tc, pc := dfs(root)
	return tc - pc
}
func minimalExecTime(root *TreeNode) float64 {
	//求每个节点及其子节点,串行的执行时间和答案
	var dfs func(node *TreeNode) (float64, float64)
	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0, 0
		}
		left1, left2 := dfs(node.Left)
		right1, right2 := dfs(node.Right)
		return float64(node.Val) + left1 + right1, float64(node.Val) + math.Max(math.Max(left2, right2), (left1+right1)/2)
	}
	_, pc := dfs(root)
	return pc
}
