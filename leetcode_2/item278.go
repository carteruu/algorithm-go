package leetcode_2

func firstBadVersion(n int) int {
	l, r := 0, n
	for l <= r {
		mid := (l + r) >> 1
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

var isBadVersion func(version int) bool
