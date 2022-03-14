package leetcode

import "algorithm/pkg"

func FindMode(root *pkg.TreeNode) []int {
	return findMode(root)
}
func findMode(root *pkg.TreeNode) []int {
	m := make(map[int]int)
	dfs(root, m)

	times := 0
	var rs []int
	for k, v := range m {
		if v == times {
			rs = append(rs, k)
		}
		if v > times {
			rs = []int{k}
			times = v
		}
	}
	return rs
}

func dfs(node *pkg.TreeNode, m map[int]int) {
	if node == nil {
		return
	}
	m[node.Val]++
	dfs(node.Left, m)
	dfs(node.Right, m)
}

func findMode1(root *pkg.TreeNode) []int {
	res := []int{}
	base, count, maxCount := 0, 0, 0
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	var inOrder func(root *pkg.TreeNode)
	inOrder = func(root *pkg.TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		update(root.Val)
		inOrder(root.Right)
	}
	inOrder(root)

	return res
}
