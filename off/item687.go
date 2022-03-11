package off

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(node *TreeNode, pre int) int
	dfs = func(node *TreeNode, pre int) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left, node.Val)
		right := dfs(node.Right, node.Val)
		//选择当前节点作为根的路径长度
		l := 1 + left + right
		if l > res {
			res = l
		}

		if node.Val != pre {
			//当前节点值和父节点值不同
			return 0
		} else {
			//当前节点值和父节点值相同,返回当前节点+子树节点
			//因为路径只能有两个终点,所以返回左右子树,数量比较多的
			if left > right {
				return 1 + left
			} else {
				return 1 + right
			}
		}
	}
	dfs(root, -1)
	return res - 1
}
