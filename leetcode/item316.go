package leetcode

func removeDuplicateLetters1(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}

func removeDuplicateLetters(s string) string {
	m := [26]int{}
	//计算每个字母的个数
	for _, c := range s {
		m[c-'a']++
	}
	//记录是否已经在栈中
	inStack := [26]bool{}
	//单调栈
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if !inStack[s[i]-'a'] {
			//当前字符不在栈中
			for size := len(stack); size > 0 && m[stack[size-1]-'a'] > 0 && stack[size-1] > s[i]; size = len(stack) {
				//栈存在字母,栈顶的字母在后面还存在,栈顶的字母大于当前字母:可以弹出
				//弹出栈顶,并设置状态
				inStack[stack[size-1]-'a'] = false
				stack = stack[:size-1]
			}
			//将当前字母压入栈,并设置状态
			stack = append(stack, s[i])
			inStack[s[i]-'a'] = true
		}
		//后面存在的字母数减1
		m[s[i]-'a']--
	}
	return string(stack)
}
