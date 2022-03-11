package off

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var isEq func(a *TreeNode, b *TreeNode) bool
	isEq = func(a *TreeNode, b *TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil {
			return false
		}

		return a.Val == b.Val && isEq(a.Left, b.Left) && isEq(a.Right, b.Right)
	}

	if A == nil || B == nil {
		return false
	}
	return isEq(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
