package leetcode_2

func countStudents(students []int, sandwiches []int) int {
	s0 := 0
	s1 := 0
	for _, s := range students {
		s0 += 1 - s
		s1 += s
	}
	for _, sandwich := range sandwiches {
		if sandwich == 1 && s1 == 0 {
			return s0
		}
		if sandwich == 0 && s0 == 0 {
			return s1
		}
		s0 -= 1 - sandwich
		s1 -= sandwich
	}
	return 0
}
