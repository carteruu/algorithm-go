package offer_1

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	l := 0
	for r := 1; r < len(nums); r++ {
		sum += nums[r]
		if sum > max {
			max = sum
		}
		for l <= r && sum < 0 {
			sum -= nums[l]
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
	ma := nums[0]
	lSum := 0
	for _, num := range nums {
		lSum = max(lSum+num, num)
		ma = max(ma, lSum)
	}
	return ma
}
