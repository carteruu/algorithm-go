package leetcode_2

func largestAltitude(gain []int) int {
	h, m := 0, 0
	for _, g := range gain {
		h += g
		m = maxInt(m, h)
	}
	return m
}
