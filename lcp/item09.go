package lcp

func minJump(jump []int) int {
	s := make([]int, len(jump))
	s[0] = jump[0]
	for i := 1; i < len(jump); i++ {
		s[i] = i + jump[i]
		if s[i-1] > s[i] {
			s[i] = s[i-1]
		}
	}
	pos1 := 0
	step1 := 0
	for {
		//跳一步
		next := pos1 + jump[pos1]
		step1++
		if next >= len(jump) {
			return step1
		}
		//跳第二步
		next += jump[next]
		pos1 = s[pos1]
		step1++
		if next > pos1 {
			pos1 = next
		}
		if pos1 >= len(jump) {
			break
		}
	}

	pos2 := jump[0]
	step2 := 1
	for {
		//跳一步
		next := pos2 + jump[pos2]
		step2++
		if next >= len(jump) {
			return step2
		}
		//跳第二步
		next += jump[next]
		pos2 = s[pos2]
		step2++
		if next > pos2 {
			pos2 = next
		}
		if pos2 >= len(jump) {
			break
		}
	}
	if step1 < step2 {
		return step1
	}
	return step2
}
