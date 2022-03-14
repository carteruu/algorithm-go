package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildByLevelInterface(arr []interface{}) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	var q []*TreeNode

	root := &TreeNode{Val: arr[0].(int)}
	cur := root
	for i := 1; i < len(arr); i++ {
		var childNode *TreeNode
		if arr[i] != nil {
			childNode = &TreeNode{Val: arr[i].(int)}
			q = append(q, childNode)
		}
		if i&1 == 1 {
			if cur != nil {
				cur.Left = childNode
			}
		} else {
			if cur != nil {
				cur.Right = childNode
			}
			cur = q[0]
			q = q[1:]
		}
	}

	return root
}
func BuildByLevel(arr []int) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}

	var q []*TreeNode
	idx := 0

	root := &TreeNode{Val: arr[0]}
	cur := root
	for i := 1; i < len(arr); i++ {
		if i&1 == 1 {
			cur.Left = &TreeNode{Val: arr[i]}
			q = append(q, cur.Left)
		} else {
			cur.Right = &TreeNode{Val: arr[i]}
			q = append(q, cur.Right)
			cur = q[idx]
			idx++
		}
	}

	return root
}
