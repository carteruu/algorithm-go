package leetcode_2

func getLucky1(s string, k int) int {
	var sum int
	for _, c := range s {
		num := int(c - 'a' + 1)
		if num >= 10 {
			sum += num / 10
		}
		sum += num % 10
	}
	for i := 0; i < k-1 && sum >= 10; i++ {
		var temp int
		for sum >= 0 {
			temp += sum % 10
			sum /= 10
		}
		sum = temp
	}
	return sum
}
func getLucky(s string, k int) int {
	nums := make([]int, 0, len(s)*2)
	for _, c := range s {
		num := int(c - 'a' + 1)
		if num >= 10 {
			nums = append(nums, num/10)
		}
		nums = append(nums, num%10)
	}
	for i := 0; i < k && len(nums) > 1; i++ {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		nums = nil
		for sum > 0 {
			nums = append(nums, sum%10)
			sum /= 10
		}
		for j := 0; j < len(nums)/2; j++ {
			nums[j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[j]
		}
	}
	ans := 0
	n := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans += nums[i] * n
		n *= 10
	}
	return ans
}
