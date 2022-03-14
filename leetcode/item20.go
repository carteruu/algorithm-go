package leetcode

func isValid1(s string) bool {
	stack := make([]byte, 0, len(s))
	m := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := m[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || m[stack[len(stack)-1]] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
