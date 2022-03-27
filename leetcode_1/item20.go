package leetcode_1

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	for _, c := range s {
		if _, ok := m[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && m[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
