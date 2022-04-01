package leetcode_1

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, v := range inorder {
		m[v] = i
	}
	var dfs func(preIdx, l, r int) (*TreeNode, int)
	dfs = func(preIdx, l, r int) (*TreeNode, int) {
		if l >= r {
			//中序为空，直接返回空和前序
			return nil, preIdx
		}
		val := preorder[preIdx]
		node := &TreeNode{Val: val}
		//这里找出当前节点在中序遍历中的下标，从而将中序遍历分成左右两部分。
		idx := m[val]
		node.Left, preIdx = dfs(preIdx+1, l, idx)
		node.Right, preIdx = dfs(preIdx, idx+1, r)
		return node, preIdx
	}
	root, _ := dfs(0, 0, len(inorder))
	return root
}

func buildTree_1(preorder []int, inorder []int) *TreeNode {
	var dfs func(pre, in []int) (*TreeNode, []int)
	dfs = func(pre, in []int) (*TreeNode, []int) {
		if len(in) == 0 {
			return nil, pre
		}
		val := pre[0]
		pre = pre[1:]
		node := &TreeNode{Val: val}
		i := 0
		for ; i < len(in); i++ {
			if in[i] == val {
				break
			}
		}
		node.Left, pre = dfs(pre, in[:i])
		if i < len(in) {
			node.Right, pre = dfs(pre, in[i+1:])
		}
		return node, pre
	}
	root, _ := dfs(preorder, inorder)
	return root
}
