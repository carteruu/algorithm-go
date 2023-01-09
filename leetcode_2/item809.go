package leetcode_2

func expressiveWords(s string, words []string) int {
	if len(s) == 0 || len(words) == 0 {
		return 0
	}
	cs := make([]byte, 0, len(s))
	cnts := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			cnts[len(cnts)-1]++
		} else {
			cs = append(cs, s[i])
			cnts = append(cnts, 1)
		}
	}

	ans := 0
	for _, word := range words {
		if word == "" {
			continue
		}
		if word[0] != cs[0] {
			continue
		}
		cnt := cnts[0] - 1
		idx := 0
		i := 1
		for ; i < len(word); i++ {
			if cnt == 0 || (word[i] != cs[idx] && cnts[idx] >= 3) {
				idx++
				if idx == len(cs) {
					break
				}
				cnt = cnts[idx]
			}
			if idx == len(cs) || word[i] != cs[idx] {
				break
			}
			cnt--
		}
		if (i == len(word) && idx == len(cs)-1) && (cnt == 0 || cnts[idx] >= 3) {
			ans++
		}
	}
	return ans
}
