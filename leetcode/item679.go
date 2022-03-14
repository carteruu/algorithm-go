package leetcode

func judgePoint24(cards []int) bool {
	var backtrack func(idx, s int) bool
	backtrack = func(idx, s int) bool {
		if idx == 4 {
			return s == 24
		}
		if backtrack(idx+1, s+cards[idx]) {
			return true
		}
		if backtrack(idx+1, s-cards[idx]) {
			return true
		}
		if backtrack(idx+1, s*cards[idx]) {
			return true
		}
		if s%cards[idx] == 0 && backtrack(idx+1, s/cards[idx]) {
			return true
		}

		if idx < 3 {
			nums := []int{cards[idx] + cards[idx+1], cards[idx] - cards[idx+1], cards[idx] * cards[idx+1]}
			if cards[idx]%cards[idx+1] == 0 {
				nums = append(nums, cards[idx]/cards[idx+1])
			}
			for _, num := range nums {
				if backtrack(idx+2, s+num) {
					return true
				}
				if backtrack(idx+2, s-num) {
					return true
				}
				if backtrack(idx+2, s*num) {
					return true
				}
				if num != 0 && s%num == 0 && backtrack(idx+2, s/num) {
					return true
				}
			}
		}
		return false
	}
	return backtrack(0, 0)
}
