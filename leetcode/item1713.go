package leetcode

import "sort"

func minOperations(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

func minOperations3(target []int, arr []int) int {
	mT := make(map[int]struct{})
	for _, num := range target {
		mT[num] = struct{}{}
	}
	mA := make(map[int]struct{})
	for _, num := range arr {
		mA[num] = struct{}{}
	}
	var targetNew []int
	var arrNew []int
	var res int
	for _, num := range target {
		if _, ok := mA[num]; ok {
			targetNew = append(targetNew, num)
		} else {
			res++
		}
	}
	for _, num := range arr {
		if _, ok := mT[num]; ok {
			arrNew = append(arrNew, num)
		}
	}

	return res + dynamic(targetNew, arrNew)
}
func dynamic(target []int, arr []int) int {
	lenA := len(arr)
	lenT := len(target)
	if lenT == 0 || lenA == 0 {
		return lenT
	}
	dp := make([]int, lenA)
	if target[0] != arr[0] {
		dp[0] = 1
	}
	for i := 1; i < lenA; i++ {
		if target[0] == arr[i] {
			dp[i] = 0
		} else {
			dp[i] = dp[i-1]
		}
	}
	for i := 1; i < lenT; i++ {
		if arr[0] != target[i] {
			dp[0]++
		}
		for j := 1; j < lenA; j++ {
			if arr[j] == target[i] {
				//匹配
				if dp[j-1] < dp[j] {
					dp[j] = dp[j-1]
				}
			} else {
				dp[j]++
				if dp[j-1] < dp[j] {
					dp[j] = dp[j-1]
				}
			}
		}
	}
	return dp[lenA-1]
}

func minOperations1(target []int, arr []int) int {
	lenT := len(target)
	lenA := len(arr)
	if lenT == 0 || lenA == 0 {
		return lenT
	}
	dp := make([][]int, lenA)
	dp[0] = make([]int, lenT)
	for i := 0; i < lenA; i++ {
		dp[i] = make([]int, lenT)
		if target[0] == arr[i] {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}
	for i := 1; i < lenT; i++ {
		if arr[0] == target[i] {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = dp[0][i-1] + 1
		}
	}
	for i := 1; i < lenA; i++ {
		for j := 1; j < lenT; j++ {
			if arr[i] == target[j] {
				//匹配
				if dp[i][j-1] < dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if dp[i][j-1]+1 < dp[i-1][j] {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[lenA-1][lenT-1]
}
