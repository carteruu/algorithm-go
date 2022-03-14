package leetcode

import "algorithm/pkg"

type BSTIterator12 struct {
	stack []*pkg.TreeNode
	cur   *pkg.TreeNode
}

func ConstructorBSTIterator12(root *pkg.TreeNode) BSTIterator12 {
	return BSTIterator12{cur: root}
}

func (it *BSTIterator12) Next() int {
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}
	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator12) HasNext() bool {
	return it.cur != nil || len(it.stack) > 0
}
