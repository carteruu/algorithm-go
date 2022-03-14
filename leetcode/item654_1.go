package leetcode

import "algorithm/pkg"

func constructMaximumBinaryTree11(nums []int) *pkg.TreeNode {
	//单调栈-递增
	stack := make([]*pkg.TreeNode, 0, len(nums))
	for _, num := range nums {
		//当前节点
		curr := &pkg.TreeNode{Val: num}
		//将栈顶比当前节点值小的元素出栈
		for size := len(stack); size > 0 && stack[size-1].Val < num; size = len(stack) {
			//最后一个小于当前节点值的元素为左子节点
			curr.Left = stack[size-1]
			stack = stack[:size-1]
		}
		if len(stack) > 0 {
			//栈不为空,则当前节点为栈顶的右子节点,
			//后面继续循环,可能又会出现比当前节点值大的元素,将会替换掉当前节点,并将当前节点作为其左子节点
			stack[len(stack)-1].Right = curr
		}
		stack = append(stack, curr)
	}
	//栈底为最大元素,即根节点
	return stack[0]
}
func constructMaximumBinaryTree12(nums []int) *pkg.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &pkg.TreeNode{Val: nums[0]}
	}
	//找到最大元素的下标,作为根节点,并分隔左右子树
	maxIdx := 0
	for i, num := range nums {
		if num > nums[maxIdx] {
			maxIdx = i
		}
	}
	//递归处理
	return &pkg.TreeNode{
		Val:   nums[maxIdx],
		Left:  constructMaximumBinaryTree(nums[:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
