package leetcode_2

func reinitializePermutation(n int) int {
	i, ans := 1, 0
	for {
		if i&1 == 0 {
			i = i / 2
		} else {
			i = n/2 + (i-1)/2
		}
		ans++
		if i == 1 {
			return ans
		}
	}
}
func reinitializePermutation1(n int) int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	ans := 0
	for {
		is := true
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
			if arr[i] != i {
				is = false
			}
		}
		ans++
		if is {
			return ans
		}
		perm = arr
	}
	return ans
}
