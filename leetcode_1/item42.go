package leetcode_1

func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		//贪心：某个位置能接到的雨水 = min(左边最大高度，右边最大高度) - 当前位置高度。
		//所以如果左边的最大值 < 右边的最大值，则计算左指针能接的雨水；否则计算右边
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
func trap_3(height []int) int {
	n := len(height)
	hh := make([]int, n, n)
	mx := 0
	for i := 0; i < n; i++ {
		mx = max(mx, height[i])
		hh[i] = mx
	}
	mx = 0
	for i := n - 1; i >= 0; i-- {
		mx = max(mx, height[i])
		hh[i] = min(hh[i], mx)
	}
	ans := 0
	for i := 0; i < n; i++ {
		s := hh[i] - height[i]
		if s > 0 {
			ans += s
		}
	}
	return ans
}
func trap_2(height []int) int {
	n := len(height)
	left := make([]int, n, n)
	right := make([]int, n, n)
	leftMax := 0
	rightMax := 0
	for i := 0; i < n; i++ {
		leftMax = max(leftMax, height[i])
		left[i] = leftMax

		rightMax = max(rightMax, height[n-1-i])
		right[n-1-i] = rightMax
	}
	ans := 0
	for i := 0; i < n; i++ {
		s := min(left[i], right[i]) - height[i]
		if s > 0 {
			ans += s
		}
	}
	return ans
}

//每一个位置能盛的水 = min(左边最高，右边最高）- 该位置高度
func trap_1(height []int) int {
	n := len(height)
	ans := 0
	for i := 1; i < n-1; i++ {
		leftMax := 0
		for j := 0; j < i; j++ {
			leftMax = max(leftMax, height[j])
		}
		rightMax := 0
		for j := n - 1; j > i; j-- {
			rightMax = max(rightMax, height[j])
		}
		s := min(leftMax, rightMax) - height[i]
		if s > 0 {
			ans += s
		}
	}
	return ans
}
