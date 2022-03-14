package leetcode

func isCovered(ranges [][]int, left, right int) bool {
	diff := [52]int{} // 差分数组
	for _, r := range ranges {
		//进出区间
		diff[r[0]]++
		diff[r[1]+1]--
	}
	cnt := 0 // 前缀和
	for i := 1; i <= 50; i++ {
		cnt += diff[i]
		if cnt <= 0 && left <= i && i <= right {
			return false
		}
	}
	return true
}

func isCovered1(ranges [][]int, left int, right int) bool {
	a := struct{}{}
	m := make(map[int]struct{})
	for _, r := range ranges {
		for num := r[0]; num <= r[1]; num++ {
			m[num] = a
		}
	}

	for num := left; num <= right; num++ {
		if _, ok := m[num]; !ok {
			return false
		}
	}
	return true
}
