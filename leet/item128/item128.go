package item128

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	ans := 0
	for num := range m {
		l := 1
		if _, ok := m[num-1]; !ok {
			//这个数是起始
			for i := num + 1; ; i++ {
				if _, ok := m[i]; !ok {
					break
				}
				l++
			}
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}
func longestConsecutive2(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	ans := 0
	for len(m) > 0 {
		var num int
		l := 1
		//随机取一个数字
		for num = range m {
			break
		}
		delete(m, num)
		for lNum := num - 1; ; lNum-- {
			if _, ok := m[lNum]; !ok {
				break
			}
			l++
			delete(m, lNum)
		}
		for rNum := num + 1; ; rNum++ {
			if _, ok := m[rNum]; !ok {
				break
			}
			l++
			delete(m, rNum)
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}

func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
