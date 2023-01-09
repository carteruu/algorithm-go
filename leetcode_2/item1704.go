package leetcode_2

func halvesAreAlike(s string) bool {
	var cnt1, cnt2 int
	set := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	for i := 0; i < len(s)/2; i++ {
		if _, ok := set[s[i]]; ok {
			cnt1++
		}
		if _, ok := set[s[i+len(s)/2]]; ok {
			cnt2++
		}
	}
	return cnt1 == cnt2
}
