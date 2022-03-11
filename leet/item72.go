package leet

func minDistance(word1 string, word2 string) int {
	var dis func(i, j int) int
	dis = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2) {
			return len(word1) - i
		}
		if word1[i] == word2[j] {
			return dis(i+1, j+1)
		} else {
			//插入
			a := dis(i, j+1) + 1
			//删除
			b := dis(i+1, j) + 1
			//替换
			c := dis(i+1, j+1) + 1
			if a > b {
				a = b
			}
			if a > c {
				a = c
			}
			return a
		}
	}
	return dis(0, 0)
}
