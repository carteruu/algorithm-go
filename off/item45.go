package off

import (
	"bytes"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	//转换成字符串,方便比较字符
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	sort.Slice(strs, func(i, j int) bool {
		li := len(strs[i])
		lj := len(strs[j])
		//循环比较每一个字符,这里最多需要遍历到两个字符串长度的和
		//如:121,12:
		//(1) 先比较下标0和1,相同;
		//(2) 再比较下标2,取模后比较下标2和0,相同;
		//(3) 继续比较下标3,取模后比较下标0和1,121的第0个字符比12的第1和字符小,所以121排前面
		for k := 0; k < li+lj; k++ {
			if strs[i][k%li] == strs[j][k%lj] {
				continue
			}
			return strs[i][k%li] < strs[j][k%lj]
		}
		return false
	})
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}
