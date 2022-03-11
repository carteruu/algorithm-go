package off

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var match func(left, right *TreeNode) bool
	match = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && match(left.Left, right.Right) && match(left.Right, right.Left)
	}
	return match(root.Left, root.Right)
}
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := make([]*TreeNode, 0, 500)
	q = append(q, root.Left, root.Right)
	for len(q) > 0 {
		size := len(q)
		isAllNil := true
		for i := 0; i < size/2; i++ {
			head := q[i]
			tail := q[size-1-i]
			isHeadNil := head == nil
			isTailNil := tail == nil
			if isHeadNil && !isTailNil {
				return false
			}
			if !isHeadNil && isTailNil {
				return false
			}
			if !isHeadNil && !isTailNil {
				if head.Val != tail.Val {
					return false
				}
			}
		}
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if poll != nil {
				q = append(q, poll.Left, poll.Right)
				isAllNil = false
			} else {
				q = append(q, nil, nil)
			}
		}
		if isAllNil {
			break
		}
	}
	return true
}
