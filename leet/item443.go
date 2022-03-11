package leet

import "strconv"

func compress(chars []byte) int {
	slow, fast := 0, 1
	cnt := 1
	for fast < len(chars) {
		if chars[fast] == chars[slow] {
			cnt++
		} else {
			if cnt > 1 {
				str := strconv.Itoa(cnt)
				for _, c := range str {
					slow++
					chars[slow] = byte(c)
				}
				cnt = 1
			}
			slow++
			chars[slow] = chars[fast]
		}
		fast++
	}
	if cnt > 1 {
		str := strconv.Itoa(cnt)
		for _, c := range str {
			slow++
			chars[slow] = byte(c)
		}
	}
	return slow + 1
}

func compress1(chars []byte) int {
	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}
			left = read + 1
		}
	}
	return write
}
