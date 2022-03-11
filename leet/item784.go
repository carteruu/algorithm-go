package leet

import (
	"bytes"
)

func letterCasePermutation(s string) []string {
	n := len(s)
	ans := make([]string, 0, 1<<n)
	for mark := 0; mark < 1<<n; mark++ {
		var buffer bytes.Buffer
		hasNum := false
		for i := 0; i < n; i++ {
			if mark&(1<<i) == 0 {
				buffer.WriteByte(s[i])
			} else {
				if s[i] >= '0' && s[i] <= '9' {
					hasNum = true
					break
				} else if s[i] >= 'a' && s[i] <= 'z' {
					buffer.WriteByte(s[i] - 32)
				} else if s[i] >= 'A' && s[i] <= 'Z' {
					buffer.WriteByte(s[i] + 32)
				}
			}
		}
		if !hasNum {
			ans = append(ans, buffer.String())
		}
	}
	return ans
}
func letterCasePermutation1(s string) []string {
	n := len(s)
	m := make(map[string]struct{}, 1<<n)
	for mark := 0; mark < 1<<n; mark++ {
		var buffer bytes.Buffer
		for i := 0; i < n; i++ {
			if mark&(1<<i) == 0 || s[i] >= '0' && s[i] <= '9' {
				buffer.WriteByte(s[i])
			} else {
				if s[i] >= 'a' && s[i] <= 'z' {
					buffer.WriteByte(s[i] - 32)
				} else if s[i] >= 'A' && s[i] <= 'Z' {
					buffer.WriteByte(s[i] + 32)
				}
			}
		}
		m[buffer.String()] = struct{}{}
	}
	ans := make([]string, 0, len(m))
	for str := range m {
		ans = append(ans, str)
	}
	return ans
}
