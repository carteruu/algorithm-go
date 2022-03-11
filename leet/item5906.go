package leet

import "strings"

func countValidWords(sentence string) int {
	stateSets := map[string]map[string]struct{}{
		"start":       {"letter": struct{}{}, "punctuation": struct{}{}},
		"letter":      {"letter": struct{}{}, "-": struct{}{}, "punctuation": struct{}{}, "end": struct{}{}},
		"num":         {},
		"-":           {"letter": struct{}{}},
		"punctuation": {"end": struct{}{}},
		"end":         {},
	}

	strs := strings.Split(sentence, " ")
	cnt := 0
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}
		state := "start"
		isMeet := true
		_cnt := 0
		for _, c := range str {
			switch {
			case 'a' <= c && c <= 'z':
				{
					if _, ok := stateSets[state]["letter"]; !ok {
						isMeet = false
						break
					}
					state = "letter"
				}
			case '0' <= c && c <= '9':
				{
					if _, ok := stateSets[state]["num"]; !ok {
						isMeet = false
						break
					}
					state = "num"
				}
			case c == '-':
				{
					if _, ok := stateSets[state]["-"]; !ok {
						isMeet = false
						break
					}
					state = "-"
					_cnt++
				}
			case c == '!' || c == ',' || c == '.':
				{
					if _, ok := stateSets[state]["punctuation"]; !ok {
						isMeet = false
						break
					}
					state = "punctuation"
				}
			}
			if !isMeet {
				break
			}
		}
		if _, ok := stateSets[state]["end"]; !isMeet || !ok || _cnt > 1 {
			continue
		}
		cnt++
	}
	return cnt
}
