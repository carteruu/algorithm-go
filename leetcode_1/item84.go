package leetcode_1

func largestRectangleArea(heights []int) int {
	ans := 0
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}
	leftStack := make([]int, 0, n)
	rightStack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for last := len(leftStack) - 1; len(leftStack) > 0 && heights[leftStack[last]] > heights[i]; last-- {
			right[leftStack[last]] = i
			leftStack = leftStack[:last]
		}
		leftStack = append(leftStack, i)

		for last := len(rightStack) - 1; len(rightStack) > 0 && heights[rightStack[last]] > heights[n-1-i]; last-- {
			left[rightStack[last]] = n - 1 - i
			rightStack = rightStack[:last]
		}
		rightStack = append(rightStack, n-1-i)
	}
	for i, height := range heights {
		ans = max(ans, (right[i]-left[i]-1)*height)
	}
	return ans
}

//找到每个下标，左右两边 >= 当前高度的最远距离
func largestRectangleArea_1(heights []int) int {
	ans := 0
	for i, height := range heights {
		l := i
		r := i
		for ; l >= 0 && heights[l] >= height; l-- {

		}
		for ; r < len(heights) && heights[r] >= height; r++ {

		}
		ans = max(ans, (r-l-1)*height)
	}
	return ans
}
