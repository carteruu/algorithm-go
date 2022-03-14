package offer

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ans := "1"
	for i := 1; i < n; i++ {
		var str bytes.Buffer
		var c = ans[0]
		var cnt = 1
		for j := 1; j < len(ans); j++ {
			if ans[j] == c {
				cnt++
			} else {
				str.WriteString(strconv.Itoa(cnt))
				str.WriteString(string(c))
				c = ans[j]
				cnt = 1
			}
		}
		str.WriteString(strconv.Itoa(cnt))
		str.WriteString(string(c))
		ans = str.String()
	}
	return ans
}
