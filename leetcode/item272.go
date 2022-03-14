package leetcode

import "algorithm/pkg"

func closestKValues(root *pkg.TreeNode, target float64, k int) []int {
	ans := make([]int, 0)
	var dfs func(node *pkg.TreeNode)
	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if len(ans) < k {
			ans = append(ans, node.Val)
		} else {
			diffStart := diff(ans[0], target)
			diffEnd := diff(ans[len(ans)-1], target)
			d := diff(node.Val, target)
			if d < diffStart || d < diffEnd {
				if diffStart > diffEnd {
					ans = ans[1:]
				} else {
					ans = ans[:len(ans)-1]
				}
				ans = append(ans, node.Val)
			} else {
				return
			}
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
func diff(num int, target float64) float64 {
	d := float64(num) - target
	if d < 0 {
		d = -d
	}
	return d
}
