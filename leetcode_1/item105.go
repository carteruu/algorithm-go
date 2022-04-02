package leetcode_1

/**
TODO 迭代的做法
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{}
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

func buildTree_3(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := preorder[0]
	preorder = preorder[1:]
	node := &TreeNode{Val: val}
	i := 0
	for ; i < len(inorder) && inorder[i] != val; i++ {
	}
	node.Left = buildTree_3(preorder[0:len(inorder[:i])], inorder[:i])
	node.Right = buildTree_3(preorder[len(inorder[:i]):], inorder[i+1:])
	return node
}
func buildTree_2(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, v := range inorder {
		m[v] = i
	}
	/**
	中序遍历只是用来区分左右子树,[l,r)
	*/
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
