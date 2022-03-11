package off

func verifyPostorder(postorder []int) bool {
	size := len(postorder)
	if size <= 1 {
		return true
	}
	root := postorder[size-1]
	//左子树边界,不包含i
	i := 0
	for postorder[i] < root {
		i++
	}
	//右子树所有元素大于根节点
	for j := i; j < size-1; j++ {
		if postorder[j] <= root {
			return false
		}
	}
	return verifyPostorder(postorder[0:i]) && verifyPostorder(postorder[i:size-1])
}
func verifyPostorder1(postorder []int) bool {
	var verify func(l, r int) bool
	verify = func(l, r int) bool {
		if r-l < 1 {
			return true
		}
		root := postorder[r]

		leftEndIdx := r - 1
		for i := l; i < r; i++ {
			if leftEndIdx == r-1 && postorder[i] > root {
				leftEndIdx = i - 1
			}
			if leftEndIdx != r-1 && postorder[i] <= root {
				return false
			}
		}
		return (leftEndIdx < l || verify(l, leftEndIdx)) && (leftEndIdx+1 >= r || verify(leftEndIdx+1, r-1))
	}
	return verify(0, len(postorder)-1)
}
