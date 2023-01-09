package leetcode_2

func reinitializePermutation(n int) int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	arr := make([]int, n)
	ans := 0
	for {
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		perm = arr
		ans++
		is := true
		for i := 0; i < n; i++ {
			if perm[i] != i {
				is = false
				break
			}
		}
		if is {
			return ans
		}
	}
	return ans
}
