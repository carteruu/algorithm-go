package leet

import "bytes"

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	//从后往前确定需要的字母下标
	needLetterIdxs := make([]int, repetition)
	needLetterIdxsIdx := repetition - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == letter {
			needLetterIdxs[needLetterIdxsIdx] = i
			if needLetterIdxsIdx == 0 {
				break
			}
			needLetterIdxsIdx--
		}
	}
	//找出剩下的字母
	var ans bytes.Buffer
	l := 0
	for i := 0; i < k-repetition; {
		minLetter := s[l]
		r := len(s) - k + repetition - 1 - i
		if needLetterIdxsIdx < len(needLetterIdxs) && needLetterIdxs[needLetterIdxsIdx] < r {
			r = needLetterIdxs[needLetterIdxsIdx]
		}
		for j := l; j <= r; j++ {
			if s[j] <= minLetter {
				minLetter = s[j]
				l = j + 1
			}
		}
		ans.WriteByte(minLetter)
		if needLetterIdxsIdx < len(needLetterIdxs) && l-1 == needLetterIdxs[needLetterIdxsIdx] {
			needLetterIdxsIdx++
		} else {
			i++
		}
	}
	return ans.String()
}
