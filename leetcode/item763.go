package leetcode

func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	var ans []int
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
func partitionLabels1(s string) []int {
	big := make([]int, 26)
	small := make(map[int32]int)
	for _, c := range s {
		big[c]++
	}
	var ans []int
	start := 0
	for i, c := range s {
		if big[c] > 0 {
			if big[c] > 1 {
				small[c] = big[c] - 1
			}
			big[c] = 0
		} else {
			if small[c] == 1 {
				small[c]--
			} else {
				small[c]--
			}
		}
		if len(small) == 0 {
			ans = append(ans, i-start+1)
			start = i + 1
		}
	}
	return ans
}
