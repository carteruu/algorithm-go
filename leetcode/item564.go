package leetcode

func nearestPalindromic(n string) string {
	//本身就是回文，改中间；不是回文，改后半段
	l := len(n)
	mid := l / 2
	b := []byte(n)

	isAllNine := l > 1
	for i := 0; i < l && isAllNine; i++ {
		isAllNine = b[i] == '9'
	}
	if isAllNine {
		ans := make([]byte, l+1)
		ans[0] = '1'
		ans[l] = '1'
		for i := 1; i < l; i++ {
			ans[i] = '0'
		}
		return string(ans)
	}

	isZeroOne := l > 1 && b[0] == '1' && (b[l-1] == '0' || b[l-1] == '1')
	for i := 1; i < l-1 && isZeroOne; i++ {
		isZeroOne = b[i] == '0'
	}
	if isZeroOne {
		b = b[1:]
		for i := 0; i < len(b); i++ {
			b[i] = '9'
		}
		return string(b)
	}

	isP := true
	for i := 0; i < mid && isP; i++ {
		isP = n[i] == n[l-1-i]
	}
	if isP {
		t := b[mid]
		if b[mid] == '0' {
			t = '1'
		} else {
			t--
		}
		if l&1 == 0 {
			b[mid] = t
			b[mid-1] = t
		} else {
			b[mid] = t
		}
		return string(b)
	}

	for i := l - 1; i >= mid; i-- {
		b[i] = b[l-1-i]
	}
	return string(b)
}
