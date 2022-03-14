package leetcode

func findSubstring1(s string, words []string) []int {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	var ans []int
	wordLen := len(words[0])
	n := len(words)
	l := wordLen * n
	for i := 0; i < len(s)-l+1; i++ {
		is := true
		tm := make(map[string]int)
		for key, val := range m {
			tm[key] = val
		}
		for j := 0; j < n; j++ {
			start := i + wordLen*j
			w := s[start : start+wordLen]
			if cnt, ok := tm[w]; ok && cnt > 0 {
				tm[w]--
			} else {
				is = false
				break
			}
		}
		if is {
			ans = append(ans, i)
		}
	}
	return ans
}
func findSubstring(s string, words []string) []int {
	wL := len(words)           // words 中单词个数
	oneWordL := len(words[0])  // 每个单词长度
	hm := make(map[string]int) // 存words中单词的map
	Allcount := 0              // 记录不同单词的个数，用于在窗口滑动过程中判断是否满足答案条件
	for i := 0; i < wL; i++ {
		if _, ok := hm[words[i]]; !ok {
			Allcount++
		}
		hm[words[i]]++
	}

	ans := []int{}
	for i := 0; i < oneWordL; i++ { // 这里参考各位大佬的思路，从0 - oneWordL的各位置开始向后滑
		Count := 0                // 和AllCount对应，这里记录滑动窗口中不同单词的个数
		tmpHm := map[string]int{} // 滑动窗口中的单词记录map
		left, right := i, i+oneWordL
		for ; right <= len(s); right += oneWordL {
			str := s[right-oneWordL : right]
			if _, ok := hm[str]; ok {
				tmpHm[str]++
				if tmpHm[str] == hm[str] { // Count++的条件是该单词出现次数在窗口tmpHm中的出现次数和hm次数一样
					Count++
				}
			} // 如果这一步出现了不相关的单词怎么办？不影响，后面在比较Count和AllCount时，还会比较right - left == oneWordL * wL 。

			for right-left > oneWordL*wL {
				str := s[left : left+oneWordL]
				if _, ok := tmpHm[str]; ok {
					tmpHm[str]--
				}
				if tmpHm[str]+1 == hm[str] { // 这个条件重要，tmpHm[str] < hm[str]不对
					Count--
				}
				left += oneWordL
			}

			if Count == Allcount && right-left == oneWordL*wL { // 两个条件判断是否满足答案
				ans = append(ans, left)
			}
		}
	}
	return ans
}
