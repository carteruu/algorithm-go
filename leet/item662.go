package leet

import "algorithm/pkg"

func widthOfBinaryTree(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	type node struct {
		*pkg.TreeNode
		idx int
	}
	queue := make([]*node, 0, 500)
	queue = append(queue, &node{root, 0})
	for len(queue) > 0 {
		size := len(queue)
		l := queue[size-1].idx - queue[0].idx + 1
		if l > ans {
			ans = l
		}
		for i := 0; i < size; i++ {
			poll := queue[0]
			queue = queue[1:]
			if poll.Left != nil {
				queue = append(queue, &node{
					poll.Left,
					poll.idx * 2,
				})
			}
			if poll.Right != nil {
				queue = append(queue, &node{
					poll.Right,
					poll.idx*2 + 1,
				})
			}
		}
	}
	return ans
}
func widthOfBinaryTree1(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	var queue []*pkg.TreeNode
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		size := len(queue)
		l := size
		isLeftFind := false
		isRightFind := false
		for i := 0; i < size; i++ {
			left := queue[0]
			right := queue[len(queue)-1]
			queue = queue[1 : len(queue)-1]
			if left != nil {
				isLeftFind = true
				queue = append(queue, left.Left, left.Right)
			} else {
				if !isLeftFind {
					l--
				}
				queue = append(queue, nil, nil)
			}
			if right != nil {
				isRightFind = true
				queue = append(queue, right.Left, right.Right)
			} else {
				if !isRightFind {
					l--
				}
				queue = append(queue, nil, nil)
			}
		}
		if l == 0 {
			break
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
