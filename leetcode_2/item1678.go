package leetcode_2

import "bytes"

func interpret(command string) string {
	s := bytes.Buffer{}
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			s.WriteByte('G')
		} else {
			if command[i+1] == ')' {
				s.WriteByte('o')
				i++
			} else {
				s.WriteString("al")
				i += 3
			}
		}
	}
	return s.String()
}
