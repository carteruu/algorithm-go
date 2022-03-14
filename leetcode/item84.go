package leetcode

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	hs := make([]int, len(heights)+2)
	for i := 0; i < len(heights); i++ {
		hs[i+1] = heights[i]
	}
	n := len(hs)

	stack := make([]int, 0, n)
	stack = append(stack, 0)
	ans := 0
	for i := 0; i < n; i++ {
		for hs[stack[len(stack)-1]] > hs[i] {
			h := hs[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			area := (i - stack[len(stack)-1] - 1) * h
			if area > ans {
				ans = area
			}
		}
		stack = append(stack, i)
	}
	return ans
}

/*
可以使用单调栈,在线性时间找出各个柱子左右两边第一个矮柱子
*/
func largestRectangleArea7(heights []int) int {
	n := len(heights)
	stackRight := make([]int, 0, n)
	minRight := make([]int, n)
	minLeft := make([]int, n)
	for i := 0; i < n; i++ {
		minRight[i] = n
		size := len(stackRight)
		for ; size > 0 && heights[stackRight[size-1]] > heights[i]; size = len(stackRight) {
			minRight[stackRight[size-1]] = i
			stackRight = stackRight[:size-1]
		}
		if size > 0 {
			minLeft[i] = heights[stackRight[size-1]]
		} else {
			minLeft[i] = -1
		}
		stackRight = append(stackRight, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		if (minRight[i]-minLeft[i]-1)*heights[i] > ans {
			ans = (minRight[i] - minLeft[i] - 1) * heights[i]
		}
	}
	return ans
}
func largestRectangleArea4(heights []int) int {
	n := len(heights)
	stackRight := make([]int, 0, n)
	stackLeft := make([]int, 0, n)
	minRight := make([]int, n)
	minLeft := make([]int, n)
	for i := 0; i < n; i++ {
		minRight[i] = n
		for size := len(stackRight); size > 0 && heights[stackRight[size-1]] > heights[i]; size = len(stackRight) {
			minRight[stackRight[size-1]] = i
			stackRight = stackRight[:size-1]
		}
		stackRight = append(stackRight, i)

		idx := n - 1 - i
		minLeft[idx] = -1
		for size := len(stackLeft); size > 0 && heights[stackLeft[size-1]] > heights[idx]; size = len(stackLeft) {
			minLeft[stackLeft[size-1]] = idx
			stackLeft = stackLeft[:size-1]
		}
		stackLeft = append(stackLeft, idx)
	}
	ans := 0
	for i := 0; i < n; i++ {
		if (minRight[i]-minLeft[i]-1)*heights[i] > ans {
			ans = (minRight[i] - minLeft[i] - 1) * heights[i]
		}
	}
	return ans
}

func largestRectangleArea5(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func largestRectangleArea3(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	stackLeft := make([]int, 0, n)
	stackRight := make([]int, 0, n)
	for i := 0; i < n; i++ {
		size := len(stackLeft)
		for ; size > 0 && heights[stackLeft[size-1]] >= heights[i]; size = len(stackLeft) {
			stackLeft = stackLeft[:size-1]
		}
		if size > 0 {
			left[i] = stackLeft[size-1]
		} else {
			left[i] = -1
		}
		stackLeft = append(stackLeft, i)

		size = len(stackRight)
		for ; size > 0 && heights[stackRight[size-1]] >= heights[n-1-i]; size = len(stackRight) {
			stackRight = stackRight[:size-1]
		}
		if size > 0 {
			right[n-1-i] = stackRight[size-1]
		} else {
			right[n-1-i] = n
		}
		stackRight = append(stackRight, n-1-i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		if (right[i]-left[i]-1)*heights[i] > ans {
			ans = (right[i] - left[i] - 1) * heights[i]
		}
	}
	return ans
}
func largestRectangleArea2(heights []int) int {
	ans := 0
	//枚举每一根柱子
	for i, height := range heights {
		//找到左右两边第一根比当前柱子矮的柱子
		l, r := i-1, i+1
		for l >= 0 && heights[l] >= height {
			l--
		}
		for r < len(heights) && heights[r] >= height {
			r++
		}
		//计算当前柱子能构成的最大矩形(不包含左右两根比当前柱子矮的柱子)
		//(r-l+1-2)*height
		if (r-l-1)*height > ans {
			ans = (r - l - 1) * height
		}
	}
	return ans
}
func largestRectangleArea1(heights []int) int {
	ans := 0
	for i, height := range heights {
		minHeight := height
		for j := i; j < len(heights); j++ {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			if (j-i+1)*minHeight > ans {
				ans = (j - i + 1) * minHeight
			}
		}
	}
	return ans
}
