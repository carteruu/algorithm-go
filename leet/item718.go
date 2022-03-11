package leet

func findLength3(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = [2]int{0, 0}
	}
	for i := 1; i <= m; i++ {
		temp := make([][2]int, n+1)
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				temp[j][0] = dp[j-1][0] + 1
			}
			temp[j][1] = max(dp[j][1], temp[j-1][1])
			temp[j][1] = max(temp[j][1], temp[j][0])
		}
		dp = temp
	}
	return dp[n][1]
}
func findLength1(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][][2]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][2]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j][0] = dp[i-1][j-1][0] + 1
			}
			dp[i][j][1] = max(dp[i-1][j][1], dp[i][j-1][1])
			dp[i][j][1] = max(dp[i][j][1], dp[i][j][0])
		}
	}
	return dp[m][n][1]
}

func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([]int, n+1)
	var temp []int
	ans := 0
	for i := 1; i <= m; i++ {
		temp = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				temp[j] = dp[j-1] + 1
			}
			if temp[j] > ans {
				ans = temp[j]
			}
		}
		dp = temp
	}
	return ans
}
func findLength2(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func findLength11(A []int, B []int) int {
	n, m := len(A), len(B)
	ret := 0
	for i := 0; i < n; i++ {
		len := min(m, n-i)
		maxLen := maxLength(A, B, i, 0, len)
		ret = max(ret, maxLen)
	}
	for i := 0; i < m; i++ {
		len := min(n, m-i)
		maxLen := maxLength(A, B, 0, i, len)
		ret = max(ret, maxLen)
	}
	return ret
}

func maxLength(A, B []int, addA, addB, len int) int {
	ret, k := 0, 0
	for i := 0; i < len; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
