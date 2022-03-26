package leetcode_1

func rotate(nums []int, k int) {
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

//整个翻转，再翻转前k个，最后翻转剩下的
func rotate_4(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 || n == 0 {
		return
	}
	reverse := func(arr []int) {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func rotate_3(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 || n == 0 {
		return
	}
	//每次向后移动一位
	for i := 0; i < k; i++ {
		temp := nums[n-1]
		copy(nums[1:], nums[:n-1])
		nums[0] = temp
	}
}
func rotate_2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 || n == 0 {
		return
	}
	temp := make([]int, k, k)
	//先把最后k个数存起来
	for i := 0; i < k; i++ {
		temp[i] = nums[n-k+i]
	}
	//将[0,n-k)的数，放入[k,n)
	copy(nums[k:], nums[0:n-k])
	//填充前k个
	copy(nums[:k], temp)
}
func rotate_1(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 || n == 0 {
		return
	}
	temp := make([]int, n, n)
	copy(temp, nums)
	for i := 0; i < n; i++ {
		nums[i] = temp[(i-k+n)%n]
	}
}
