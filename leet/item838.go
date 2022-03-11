package leet

type pair struct {
	idx int
	c   rune
}

func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' { // 找到一段连续的没有被推动的骨牌
			j++
		}
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right { // 方向相同，那么这些竖立骨牌也会倒向同一方向
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' { // 方向相对，那么就从两侧向中间倒
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(s)
}
func pushDominoes1(dominoes string) string {
	n := len(dominoes)
	times := make([]int, n)
	cs := []rune(dominoes)
	q := make([]*pair, 0, n)
	for i, c := range cs {
		if c != '.' {
			q = append(q, &pair{i, c})
		}
	}
	t := 0
	for len(q) > 0 {
		t++
		size := len(q)
		for i := 0; i < size; i++ {
			poll := q[0]
			q = q[1:]
			if poll.c == 'R' {
				next := poll.idx + 1
				if next >= n {
					continue
				}
				if cs[next] == '.' {
					cs[next] = 'R'
					times[next] = t
					q = append(q, &pair{next, 'R'})
				} else if cs[next] == 'L' && times[next] == t {
					cs[next] = '.'
				}
			} else {
				next := poll.idx - 1
				if next < 0 {
					continue
				}
				if cs[next] == '.' {
					cs[next] = 'L'
					times[next] = t
					q = append(q, &pair{next, 'L'})
				} else if cs[next] == 'R' && times[next] == t {
					cs[next] = '.'
				}
			}
		}
	}
	return string(cs)
}
