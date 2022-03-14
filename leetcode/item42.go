package leetcode

func trap1(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		if leftMax < rightMax {
			d := leftMax - height[left+1]
			if d > 0 {
				res += d
			}
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			}
		} else {
			d := rightMax - height[right-1]
			if d > 0 {
				res += d
			}
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			}
		}
	}
	return res
}

/**
从左到右遍历数组，遍历到下标 ii 时，如果栈内至少有两个元素，记栈顶元素为 \textit{top}top，\textit{top}top 的下面一个元素是 \textit{left}left，则一定有 \textit{height}[\textit{left}] \ge \textit{height}[\textit{top}]height[left]≥height[top]。如果 \textit{height}[i]>\textit{height}[\textit{top}]height[i]>height[top]，则得到一个可以接雨水的区域，该区域的宽度是 i-\textit{left}-1i−left−1，高度是 \min(\textit{height}[\textit{left}],\textit{height}[i])-\textit{height}[\textit{top}]min(height[left],height[i])−height[top]，根据宽度和高度即可计算得到该区域能接的雨水量。
*/
func trap2(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}
