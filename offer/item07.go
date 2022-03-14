package offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for idx, val := range inorder {
		m[val] = idx
	}

	var build func(preorder []int, preStart int, preEnd int, inStart int) *TreeNode
	build = func(preorder []int, preStart int, preEnd int, inStart int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		//用中序遍历的结果，确定左右子树的长度
		inIdx := m[preorder[preStart]]
		leftLen := inIdx - inStart
		return &TreeNode{
			//用前序遍历的结果，确定父节点
			Val:   preorder[preStart],
			Left:  build(preorder, preStart+1, preStart+leftLen, inStart),
			Right: build(preorder, preStart+leftLen+1, preEnd, inIdx+1),
		}
	}
	node := build(preorder, 0, len(preorder)-1, 0)
	return node
}
