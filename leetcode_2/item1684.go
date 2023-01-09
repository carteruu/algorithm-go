package leetcode_2

func countConsistentStrings(allowed string, words []string) int {
	charSet := [26]bool{}
	for _, c := range allowed {
		charSet[c-'a'] = true
	}
	ans := len(words)
	for _, word := range words {
		for _, c := range word {
			if !charSet[c-'a'] {
				ans--
				break
			}
		}
	}
	return ans
}
