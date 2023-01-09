package leetcode_2

import "bytes"

func mergeAlternately(word1 string, word2 string) string {
	bs := bytes.Buffer{}
	i := 0
	for ; i/2 < len(word1) && (i+1)/2 < len(word2); i++ {
		if i&1 == 0 {
			bs.WriteByte(word1[i/2])
		} else {
			bs.WriteByte(word2[i/2])
		}
	}
	if len(word1) >= i/2 {
		bs.WriteString(word1[i/2:])
	}
	if len(word2) >= (i+1)/2 {
		bs.WriteString(word2[(i+1)/2:])
	}
	return bs.String()
}

func mergeAlternately1(word1 string, word2 string) string {
	bs := bytes.Buffer{}
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		bs.WriteByte(word1[i])
		bs.WriteByte(word2[i])
	}
	if len(word1) > i {
		bs.WriteString(word1[i:])
	}
	if len(word2) > i {
		bs.WriteString(word2[i:])
	}
	return bs.String()
}
