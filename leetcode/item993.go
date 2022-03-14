package leetcode

import "algorithm/pkg"

func isCousins(root *pkg.TreeNode, x, y int) bool {
	var xParent, yParent *pkg.TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	// 用来判断是否遍历到 x 或 y 的辅助函数
	update := func(node, parent *pkg.TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
	}

	type pair struct {
		node  *pkg.TreeNode
		depth int
	}
	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 && (!xFound || !yFound) {
		node, depth := q[0].node, q[0].depth
		q = q[1:]
		if node.Left != nil {
			q = append(q, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			q = append(q, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}
	}

	return xDepth == yDepth && xParent != yParent
}

func isCousins2(root *pkg.TreeNode, x, y int) bool {
	var xParent, yParent *pkg.TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node, parent *pkg.TreeNode, depth int)
	dfs = func(node, parent *pkg.TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}

		// 如果两个节点都找到了，就可以提前退出遍历
		// 即使不提前退出，对最坏情况下的时间复杂度也不会有影响
		if xFound && yFound {
			return
		}

		dfs(node.Left, node, depth+1)

		if xFound && yFound {
			return
		}

		dfs(node.Right, node, depth+1)
	}
	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}
func isCousins1(root *pkg.TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}
	xLevel := -1
	yLevel := -1

	xParent := -1
	yParent := -1

	queue := make([]*pkg.TreeNode, 0, 50)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			poll := queue[0]
			queue = queue[1:]

			if poll.Left != nil {
				if poll.Left.Val == x {
					xLevel = level + 1
					xParent = poll.Val
				}
				if poll.Left.Val == y {
					yLevel = level + 1
					yParent = poll.Val
				}
				queue = append(queue, poll.Left)
			}

			if poll.Right != nil {
				if poll.Right.Val == x {
					xLevel = level + 1
					xParent = poll.Val
				}
				if poll.Right.Val == y {
					yLevel = level + 1
					yParent = poll.Val
				}
				queue = append(queue, poll.Right)
			}
		}
		level++
	}
	return xParent != yParent && xLevel == yLevel
}
