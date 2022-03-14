package leetcode

func canCompleteCircuit3(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for idx := 0; idx < n; {
		g := 0
		i := 0
		for ; i < n; i++ {
			step := (idx + i) % n
			g += gas[step] - cost[step]
			if g < 0 {
				break
			}
		}
		if g >= 0 {
			return idx
		} else {
			idx += i + 1
		}
	}
	return -1
}
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		step := i
		g := 0
		for j := 0; j < n; j++ {
			g += gas[step] - cost[step]
			if g < 0 {
				break
			}
			step++
			step %= n
		}
		if g >= 0 {
			return i
		}
	}
	return -1
}
func canCompleteCircuit1(gas []int, cost []int) int {
	n := len(gas)
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		step := i
		g := 0
		for j := 0; j < n; j++ {
			if val, ok := m[step]; ok {
				if val[0] >= i {
					return i
				} else {
					if g < val[1] {
						return -1
					}
				}
			}
			g += gas[step] - cost[step]
			if g < 0 {
				m[i] = []int{step, g}
				break
			}
			step++
			step %= n
		}
		if g >= 0 {
			return i
		}
	}
	return -1
}
