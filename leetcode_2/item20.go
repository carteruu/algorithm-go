package leetcode_2

func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}
	m := map[int32]int32{'(': ')', '[': ']', '{': '}'}
	stack := make([]int32, 0, len(s))
	for _, c := range s {
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
