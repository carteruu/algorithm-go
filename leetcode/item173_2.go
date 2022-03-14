package leetcode

import "algorithm/pkg"

type BSTIterator struct {
	curr  *pkg.TreeNode
	stack []*pkg.TreeNode
}

func ConstructorBSTIterator(root *pkg.TreeNode) BSTIterator {
	return BSTIterator{
		curr: root,
	}
}

func (this *BSTIterator) Next() int {
	for ; this.curr != nil; this.curr = this.curr.Left {
		this.stack = append(this.stack, this.curr)
	}
	this.curr = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	val := this.curr.Val
	this.curr = this.curr.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.curr != nil || len(this.stack) > 0
}
