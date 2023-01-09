package leetcode_2

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	ans := make([]byte, maxInt(len(a), len(b))+1)
	i, j, k := len(a)-1, len(b)-1, len(ans)-1
	var s byte
	for i >= 0 || j >= 0 {
		var an, bn byte
		if i >= 0 {
			an = a[i] - '0'
		}
		if j >= 0 {
			bn = b[j] - '0'
		}
		n := an + bn + s
		ans[k] = n%2 + '0'
		s = n / 2
		i--
		j--
		k--
	}
	if s > 0 {
		ans[0] = s + '0'
		return string(ans)
	} else {
		return string(ans[1:])
	}
}

//a和b的位数很多，这种方法会溢出
func addBinary1(a string, b string) string {
	an, err := strconv.ParseInt(a, 2, 32)
	if err != nil {
		panic("")
	}
	bn, err := strconv.ParseInt(b, 2, 32)
	if err != nil {
		panic("")
	}
	for bn > 0 {
		an, bn = an^bn, (an&bn)<<1
	}
	return fmt.Sprintf("%b", an)
}
