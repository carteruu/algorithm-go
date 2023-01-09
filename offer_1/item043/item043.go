package offer_1

type CBTInserter struct {
	vals []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	c := CBTInserter{
		vals: make([]*TreeNode, 1, 1024),
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		c.vals = append(c.vals, n)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return c
}

func (r *CBTInserter) Insert(v int) int {
	child := &TreeNode{Val: v}
	r.vals = append(r.vals, child)
	parent := r.vals[(len(r.vals)-1)>>1]
	if parent.Left == nil {
		parent.Left = child
	} else {
		parent.Right = child
	}
	return parent.Val
}

func (r *CBTInserter) Get_root() *TreeNode {
	return r.vals[1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
