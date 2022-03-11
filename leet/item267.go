package leet

import "bytes"

func generatePalindromes2(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		//统计字符个数
		m[s[i]]++
	}
	//存一半的字符
	var ss []byte
	//个数为单数的字符,只允许出现一个
	singleChar := ""
	for c, val := range m {
		if val&1 == 1 {
			if singleChar == "" {
				singleChar = string(c)
				m[c]--
				val--
			} else {
				//出现多个单数字符,不能排列回文字符串
				return []string{}
			}
		}
		for i := 0; i < val; i += 2 {
			ss = append(ss, c)
		}
	}
	n := len(ss)
	var ans []string
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n-1 {
			var buffer bytes.Buffer
			for i := 0; i < n; i++ {
				buffer.WriteByte(ss[i])
			}
			buffer.WriteString(singleChar)
			for i := n - 1; i >= 0; i-- {
				buffer.WriteByte(ss[i])
			}
			ans = append(ans, buffer.String())
			return
		}
		//交换不同位置的字符
		seen := map[byte]bool{}
		for i := idx; i < n; i++ {
			if v := seen[ss[i]]; v { // 去重
				continue
			}
			ss[idx], ss[i] = ss[i], ss[idx]
			backtrack(idx + 1)
			ss[idx], ss[i] = ss[i], ss[idx]
			seen[ss[i]] = true
		}
	}
	backtrack(0)
	return ans
}

func generatePalindromes(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	var ss []byte
	singleChar := ""
	for c, val := range m {
		if val&1 == 1 {
			if singleChar == "" {
				singleChar = string(c)
				m[c]--
				val--
			} else {
				return []string{}
			}
		}
		for i := 0; i < val; i += 2 {
			ss = append(ss, c)
		}
	}
	n := len(ss)
	var ans []string
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n-1 {
			var buffer bytes.Buffer
			for i := 0; i < n; i++ {
				buffer.WriteByte(ss[i])
			}
			buffer.WriteString(singleChar)
			for i := n - 1; i >= 0; i-- {
				buffer.WriteByte(ss[i])
			}
			ans = append(ans, buffer.String())
			return
		}
		backtrack(idx + 1)
		for i := idx + 1; i < n; i++ {
			if ss[idx] == ss[i] || ss[i] == ss[i-1] {
				continue
			}
			ss[idx], ss[i] = ss[i], ss[idx]
			backtrack(idx + 1)
			ss[idx], ss[i] = ss[i], ss[idx]
		}
	}
	backtrack(0)
	return ans
}
func generatePalindromes1(s string) []string {

	cnts := map[byte]int{}
	for _, c := range s {
		cnts[byte(c)]++
	}

	bytes := []byte{}
	var odd byte
	for c, cnt := range cnts {
		if cnt%2 == 0 {
			for i := 0; i < cnt/2; i++ {
				bytes = append(bytes, c)
			}
		} else {
			if odd > 0 { // 2个以上的奇数
				return []string{}
			} else {
				for i := 0; i < (cnt-1)/2; i++ {
					bytes = append(bytes, c)
				}
				odd = c
			}
		}
	}
	n := len(bytes)

	if n == 0 {
		if odd > 0 {
			return []string{string([]byte{odd})}
		} else {
			return []string{""}
		}
	}

	// 产生所有的排列
	var ans []string
	var perm func(pos int)
	perm = func(pos int) {
		if pos == n-1 {
			tmp := make([]byte, 0, 2*n+1)
			for i := 0; i < n; i++ {
				tmp = append(tmp, bytes[i])
			}
			if odd > 0 {
				tmp = append(tmp, odd)
			}
			for i := n - 1; i >= 0; i-- {
				tmp = append(tmp, bytes[i])
			}
			ans = append(ans, string(tmp))
			return
		}
		seen := map[byte]bool{}
		for i := pos; i < n; i++ {
			if v := seen[bytes[i]]; v { // 去重
				continue
			}
			bytes[pos], bytes[i] = bytes[i], bytes[pos]
			perm(pos + 1)
			bytes[i], bytes[pos] = bytes[pos], bytes[i]
			seen[bytes[i]] = true
		}
	}

	perm(0)

	return ans
}
