package leet

import (
	"bytes"
	"strconv"
)

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
func convertToBase7_1(num int) string {
	if num == 0 {
		return "0"
	}
	abs := num
	if num < 0 {
		abs = -num
	}
	var buffer bytes.Buffer
	for abs > 0 {
		buffer.WriteByte(byte(abs%7) + '0')
		abs /= 7
	}
	b := buffer.Bytes()
	for l := 0; l < len(b)/2; l++ {
		r := len(b) - 1 - l
		b[l], b[r] = b[r], b[l]
	}
	if num < 0 {
		return "-" + string(b)
	} else {
		return string(b)
	}
}
