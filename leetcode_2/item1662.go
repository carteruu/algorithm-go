package leetcode_2

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	//i 是数组下标，j 是字符串下标
	var i1, j1 int
	var i2, j2 int
	for {
		//如果当前字符串结束
		if j1 == len(word1[i1]) {
			//如果数组也结束，则结束遍历
			if i1 == len(word1)-1 {
				break
			}
			//否则指针移动到下一个字符串的开头
			i1++
			j1 = 0
		}
		//如果当前字符串结束
		if j2 == len(word2[i2]) {
			//如果数组也结束，则结束遍历。此时第一个数组未结束，所以不相等
			if i2 == len(word2)-1 {
				return false
			}
			//否则指针移动到下一个字符串的开头
			i2++
			j2 = 0
		}
		if word1[i1][j1] != word2[i2][j2] {
			return false
		}
		//下一个字符
		j1++
		j2++
	}
	//判断第二个字符串数组是否也遍历到最后
	return j2 == len(word2[i2]) && i2 == len(word2)-1
}
