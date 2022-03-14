package item2027

func minimumMoves(s string) int {
	cnt := 0
	n := len(s)
	i := 0
	for i < n-2 {
		if s[i] == 'O' {
			i++
		} else {
			cnt++
			i += 3
		}
	}
	if (i < n && s[i] == 'X') || (i+1 < n && s[i+1] == 'X') {
		cnt++
	}
	return cnt
}
