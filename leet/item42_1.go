package leet

func trap(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax := height[l], height[r]
	ans := 0
	for l < r {
		if lMax < rMax {
			l++
			if height[l] > lMax {
				lMax = height[l]
			} else {
				ans += lMax - height[l]
			}
		} else {
			r--
			if height[r] > rMax {
				rMax = height[r]
			} else {
				ans += rMax - height[r]
			}
		}
	}
	return ans
}
