package offer

import "math"

func verifyPostorder(postorder []int) bool {
	stack := make([]int, 0, len(postorder))
	root := math.MaxInt64
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for size := len(stack); size > 0 && stack[size-1] > postorder[i]; size = len(stack) {
			root = stack[size-1]
			stack = stack[:size-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
func verifyPostorder1(postorder []int) bool {
	size := len(postorder)
	if size <= 1 {
		//0个或1个节点,返回true
		return true
	}
	//后序遍历,最后一个节点为根节点
	root := postorder[size-1]
	//从左到右,第一个大于根节点的节点之前都为左子树
	leftEnd := 0
	for ; leftEnd < size-1 && postorder[leftEnd] < root; leftEnd++ {
	}
	//遍历剩下的节点(右子树),如果有小于根节点的,不是二叉搜索树
	for i := leftEnd + 1; i < size-1; i++ {
		if postorder[i] < root {
			return false
		}
	}
	//递归判断左子树和右子树
	return verifyPostorder(postorder[:leftEnd]) && verifyPostorder(postorder[leftEnd:size-1])
}
