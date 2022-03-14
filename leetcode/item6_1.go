package leetcode

import "bytes"

func convertItem6(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := bytes.Buffer{}
	cycle := numRows*2 - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += cycle {
			b.WriteByte(s[j])
			//或 j+cycle-i*2
			second := j + (numRows-1-i)*2
			if i != 0 && i != numRows-1 && second < len(s) {
				b.WriteByte(s[second])
			}
		}
	}
	return b.String()
}
func convertItem6_1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bs := make([]bytes.Buffer, numRows)

	for i := 0; i < len(s); {
		for j := 0; j < numRows && i < len(s); j++ {
			bs[j].WriteByte(s[i])
			i++
		}
		for j := numRows - 2; j > 0 && i < len(s); j-- {
			bs[j].WriteByte(s[i])
			i++
		}
	}
	for j := 1; j < numRows; j++ {
		bs[0].WriteString(bs[j].String())
	}
	return bs[0].String()
}
