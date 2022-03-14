package leetcode

import "algorithm/pkg"

func bstFromPreorder(preorder []int) *pkg.TreeNode {
	var buildBst func(arr []int) *pkg.TreeNode
	buildBst = func(arr []int) *pkg.TreeNode {
		if len(arr) == 0 {
			return nil
		}
		leftArr := make([]int, 0, len(arr)/2)
		rightArr := make([]int, 0, len(arr)/2)
		for i := 1; i < len(arr); i++ {
			if arr[i] > arr[0] {
				rightArr = append(rightArr, arr[i])
			} else {
				leftArr = append(leftArr, arr[i])
			}
		}
		return &pkg.TreeNode{
			arr[0],
			buildBst(leftArr),
			buildBst(rightArr),
		}
	}
	return buildBst(preorder)
}
