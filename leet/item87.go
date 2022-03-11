package leet

var cache = make(map[string]bool)

func isScramble11(s1 string, s2 string) bool {
	n := len(s1)
	//字符串相同时,返回true
	if s1 == s2 {
		return true
	}
	//剩下一个字符串且不相同时,返回false
	if n == 1 && s1 != s2 {
		return false
	}
	key := s1 + "-" + s2
	if val, ok := cache[key]; ok {
		return val
	}
	//存储字符出现的次数
	m1 := make(map[byte]int, n)
	m2 := make(map[byte]int, n)
	//寻找可行的分隔点
	for i := 0; i < n-1; i++ {
		//从左向右遍历两个字符串的字符
		c1 := s1[i]
		c2 := s2[i]
		m1[c1]++
		m1[c2]--
		if m1[c1] == 0 {
			delete(m1, c1)
		}
		if m1[c2] == 0 {
			delete(m1, c2)
		}
		//当两个字符串出现的字符数量相同时,回溯并保存计算结果
		if len(m1) == 0 && isScramble(s1[:i+1], s2[:i+1]) && isScramble(s1[i+1:], s2[i+1:]) {
			cache[key] = true
			return true
		}

		//从左向右遍历是s1,从右向左遍历s2
		idx := n - 1 - i
		c3 := s2[idx]
		m2[c1]++
		m2[c3]--
		if m2[c1] == 0 {
			delete(m2, c1)
		}
		if m2[c3] == 0 {
			delete(m2, c3)
		}
		//当两个字符串出现的字符数量相同时,回溯并保存计算结果
		if len(m2) == 0 && isScramble(s1[:i+1], s2[idx:n]) && isScramble(s1[i+1:], s2[:idx]) {
			cache[key] = true
			return true
		}
	}
	//没有可行的分割点,返回false
	cache[key] = false
	return false
}

func isScramble(s1, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	// 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
	// 和谐返回 1，不和谐返回 0
	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		// 判断两个子串是否相等
		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}

		// 判断是否存在字符 c 在两个子串中出现的次数不同
		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq[:] {
			if f != 0 {
				return 0
			}
		}

		// 枚举分割位置
		for i := 1; i < length; i++ {
			// 不交换的情况
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}
			// 交换的情况
			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}

		return 0
	}
	return dfs(0, 0, n) == 1
}
