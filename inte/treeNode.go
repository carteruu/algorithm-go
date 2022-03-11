package inte

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
	q = append(q, root)
	i := 1
	for i < len(arr) {
		node := q[0]
		q = q[1:]
		if arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			q = append(q, node.Left)
		}
		i++
		if i < len(arr) {
			if arr[i] != nil {
				node.Right = &TreeNode{Val: arr[i].(int)}
				q = append(q, node.Right)
			}
		}
		i++
	}
	return root
}
