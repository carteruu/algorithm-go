package leetcode_2

import "bytes"

func largestMerge(word1 string, word2 string) string {
	var b bytes.Buffer
	for i, j := 0, 0; i < len(word1) && j < len(word2); {
		if i < len(word1) && word1[i:] > word2[j:] {
			b.WriteByte(word1[i])
			i++
		} else {
			b.WriteByte(word2[j])
			j++
		}
	}
	return b.String()
}
