package off

import "bytes"

func replaceSpace(s string) string {
	var buffer bytes.Buffer
	for _, c := range s {
		if c == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}
