package leetcode

func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return
	}

	d, t := nums[0]-nums[1], 0
	// 因为等差数列的长度至少为 3，所以可以从 i=2 开始枚举
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return
}

func numberOfArithmeticSlices2(nums []int) int {
	dp := 0
	sums := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp += 1
			sums += dp
		} else {
			dp = 0
		}
	}
	return sums
}
func numberOfArithmeticSlices1(nums []int) int {
	n := len(nums)
	res := 0
	left, right := 0, 0
	for right < n {
		if right-left < 2 {
			right++
		} else {
			if nums[left+1]-nums[left] == nums[right]-nums[right-1] {
				right++
			} else {
				res += count3(right, left)
				left = right - 1
			}
		}
	}
	return res + count3(right, left)
}

func count3(right int, left int) int {
	count := right - left
	if count > 2 {
		return (count*count - 3*count + 2) / 2
	}
	return 0
}
