package leet

import "algorithm/pkg"

func constructMaximumBinaryTree(nums []int) *pkg.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx := 0
	for i, num := range nums {
		if num > nums[maxIdx] {
			maxIdx = i
		}
	}
	return &pkg.TreeNode{
		nums[maxIdx],
		constructMaximumBinaryTree(nums[:maxIdx]),
		constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
func constructMaximumBinaryTree1(nums []int) *pkg.TreeNode {
	var build func(start, end int) *pkg.TreeNode
	build = func(start, end int) *pkg.TreeNode {
		if start > end {
			return nil
		}
		maxIdx := start
		for i := start + 1; i <= end; i++ {
			if nums[i] > nums[maxIdx] {
				maxIdx = i
			}
		}
		return &pkg.TreeNode{
			nums[maxIdx],
			build(start, maxIdx-1),
			build(maxIdx+1, end),
		}
	}
	return build(0, len(nums)-1)
}
