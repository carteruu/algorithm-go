package leetcode

import "algorithm/pkg"

type BSTIterator11 struct {
	values []int
	idx    int
}

func ConstructorBSTIterator11(root *pkg.TreeNode) BSTIterator11 {
	BSTIterator11 := BSTIterator11{}
	var dfs func(node *pkg.TreeNode)
	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		BSTIterator11.values = append(BSTIterator11.values, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return BSTIterator11
}

func (this *BSTIterator11) Next() int {
	val := this.values[this.idx]
	this.idx++
	return val
}

func (this *BSTIterator11) HasNext() bool {
	return this.idx < len(this.values)
}
