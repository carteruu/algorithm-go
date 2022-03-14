package leetcode

func generateParenthesis1(n int) []string {
	res := make([]string, 0, n^2)

	var helper func(path string, leftCount, rightCount int)
	helper = func(path string, leftCount, rightCount int) {
		if leftCount < rightCount || leftCount > n || rightCount > n {
			return
		}
		if leftCount == n && rightCount == n {
			res = append(res, path)
			return
		}
		helper(path+"(", leftCount+1, rightCount)
		helper(path+")", leftCount, rightCount+1)
	}
	helper("", 0, 0)
	return res
}
