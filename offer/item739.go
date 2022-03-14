package offer

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}
func dailyTemperatures2(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		day := 1
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			day += res[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
		}
		res[i] = day
		if len(stack) > 0 {
			res[i] = day
		} else {
			res[i] = 0
		}
		stack = append(stack, i)
	}
	return res
}
func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		day := 0
		for len(stack) > 0 && stack[len(stack)-1] < temperatures[i] {
			stack = stack[:len(stack)-1]
			day++
		}
		res[i] = day
		stack = append(stack, temperatures[i])
	}
	return res
}
