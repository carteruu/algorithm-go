package leetcode

func subArrayRanges(nums []int) int64 {
	var ans int64 = 0
	n := len(nums)
	for i := 2; i <= n; i++ {
		min := make([]int, 0, n)
		max := make([]int, 0, n)
		for j := 0; j < n; j++ {
			for len(min) > 0 && min[len(min)-1] > nums[j] {
				min = min[:len(min)-1]
			}
			min = append(min, nums[j])

			for len(max) > 0 && max[len(max)-1] < nums[j] {
				max = max[:len(max)-1]
			}
			max = append(max, nums[j])
			if j >= i-1 {
				ans += int64(max[0] - min[0])
				if min[0] == nums[j-i+1] {
					min = min[1:]
				}
				if max[0] == nums[j-i+1] {
					max = max[1:]
				}
			}
		}
	}
	return ans
}

func subArrayRanges1(nums []int) int64 {
	n := len(nums)
	var ans int64 = 0
	for i := 0; i < n; i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}
			ans += int64(max - min)
		}
	}
	return ans
}
