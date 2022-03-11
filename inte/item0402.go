package inte

func sortedArrayToBST(nums []int) *TreeNode {
	var bst func(start, end int) *TreeNode
	bst = func(start, end int) *TreeNode {
		if end < start {
			return nil
		}
		mid := (start + end) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  bst(start, mid-1),
			Right: bst(mid+1, end),
		}
	}
	return bst(0, len(nums)-1)
}
