package leetcode_2

func parseBoolExpr(expression string) bool {
	stack := make([]byte, 0, len(expression))
	for i := len(expression) - 1; i >= 0; i-- {
		if expression[i] == '(' {
			i--
			op := expression[i]
			rs := op == '&'
			for {
				c := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if c == ')' {
					if rs {
						stack = append(stack, 't')
					} else {
						stack = append(stack, 'f')
					}
					break
				}
				switch op {
				case '&':
					rs = rs && (c == 't')
				case '|':
					rs = rs || (c == 't')
				case '!':
					rs = c != 't'
				}
			}
		} else if expression[i] == ',' {
			continue
		} else {
			stack = append(stack, expression[i])
		}
	}
	return stack[0] == 't'
}
