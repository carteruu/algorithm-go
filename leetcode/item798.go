package leetcode

func bestRotation(nums []int) int {
	n := len(nums)
	maxScore := 0
	ansK := 0
	cs := make([]int, n+1)
	for i, num := range nums {
		a := (i - (n - 1) + n) % n
		b := (i - num + n) % n
		if a <= b {
			cs[a]++
			cs[b+1]--
		} else {
			cs[0]++
			cs[b+1]--
			cs[a]++
			cs[n]--
		}
	}
	score := 0
	for i, c := range cs {
		score += c
		if score > maxScore {
			maxScore = score
			ansK = i
		}
	}
	return ansK
}

//超时
func bestRotation1(nums []int) int {
	maxScore := 0
	ansK := 0
	n := len(nums)
	for k := 0; k < n; k++ {
		score := 0
		for i := 0; i < n; i++ {
			idx := (i - k + n) % n
			if nums[i] <= idx {
				score++
			}
		}
		if score > maxScore {
			maxScore = score
			ansK = k
		}
	}
	return ansK
}
