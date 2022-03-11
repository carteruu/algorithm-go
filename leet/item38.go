package leet

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	//初始值
	ans := "1"
	for i := 1; i < n; i++ {
		//第一个字符
		c := ans[0]
		cnt := 1
		//使用Buffer提高字符串拼接效率
		var buffer bytes.Buffer
		//从第二个字符开始
		for j := 1; j < len(ans); j++ {
			if ans[j] == c {
				//和上一个字符相同,数量+1
				cnt++
			} else {
				//和上一次字符不同,将上一个字符加入答案,换上新字符
				buffer.WriteString(strconv.Itoa(cnt))
				buffer.WriteByte(c)
				c = ans[j]
				cnt = 1
			}
		}
		//不要忘记拼接最后的字符
		buffer.WriteString(strconv.Itoa(cnt))
		buffer.WriteByte(c)
		ans = buffer.String()
	}
	return ans
}
