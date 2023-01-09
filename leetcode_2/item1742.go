package leetcode_2

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		num := i
		f := 0
		for num > 0 {
			f += num % 10
			num /= 10
		}
		m[f]++
	}
	max := 0
	for _, cnt := range m {
		if cnt > max {
			max = cnt
		}
	}
	return max
}
