package leetcode

func pathInZigZagTree(label int) []int {
	temp := label
	level := 1
	levelCount := 1
	for temp > levelCount {
		temp -= levelCount
		level++
		levelCount <<= 1
	}
	res := make([]int, 0, level)
	if level&1 == 0 {
		label = 1<<(level-1) + 1<<level - 1 - label
	}
	for i := level; i >= 1; i-- {
		if i&1 == 0 {
			res = append(res, 1<<(i-1)+1<<i-1-label)
		} else {
			res = append(res, label)
		}
		label >>= 1
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
