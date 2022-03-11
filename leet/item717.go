package leet

func isOneBitCharacter(bits []int) bool {
	idx := 0
	for idx < len(bits)-1 {
		if bits[idx] == 0 {
			idx++
		} else {
			idx += 2
		}
	}
	return idx == len(bits)-1 && bits[len(bits)-1] == 0
}
func isOneBitCharacter1(bits []int) bool {
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(bits)-1 {
			return bits[idx] == 0
		}
		if idx >= len(bits) {
			return false
		}
		if bits[idx] == 0 {
			return dfs(idx + 1)
		} else {
			return dfs(idx + 2)
		}
	}
	return dfs(0)
}
