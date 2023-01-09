package leetcode_2

func checkIfPangram(sentence string) bool {
	var cnts [26]bool
	cnt := 0
	for _, c := range sentence {
		if !cnts[c-'a'] {
			cnt++
			cnts[c-'a'] = true
		}
	}
	return cnt == 26
}
