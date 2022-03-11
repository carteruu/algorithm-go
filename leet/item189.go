package leet

func rotate189(nums []int, k int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	k %= n
	if k == 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
	for i := 0; i < (n-k)/2; i++ {
		nums[k+i], nums[n-1-i] = nums[n-1-i], nums[k+i]
	}
}
func rotate189_2(nums []int, k int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	k %= n
	if k == 0 {
		return
	}
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[n-k+i]
	}
	for i := 0; i < n; i++ {
		nums[i], temp[i%k] = temp[i%k], nums[i]
	}
}
func rotate189_1(nums []int, k int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	k %= n
	if k == 0 {
		return
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	for i, num := range temp {
		nums[(i+k)%len(nums)] = num
	}
}

func rotate189_3(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
