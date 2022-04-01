package leetcode_1

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjs := make([][]int, numCourses)
	ins := make([]int, numCourses)
	for _, p := range prerequisites {
		adjs[p[1]] = append(adjs[p[1]], p[0])
		ins[p[0]]++
	}
	var q []int
	for i, cnt := range ins {
		if cnt == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0, numCourses)
	for len(q) > 0 {
		poll := q[len(q)-1]
		q = q[:len(q)-1]
		ans = append(ans, poll)
		for _, next := range adjs[poll] {
			ins[next]--
			if ins[next] == 0 {
				q = append(q, next)
			}
		}
	}
	if len(ans) == numCourses {
		return ans
	}
	return nil
}
